package pluggable

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/hyperledger/fabric/integration/nwo"
	"github.com/hyperledger/fabric/integration/nwo/commands"
	"github.com/hyperledger/fabric/integration/nwo/fabricconfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/tedsuo/ifrit"
)

var _ = Describe("EndToEnd", func() {
	var (
		testDir   string
		client    *docker.Client
		network   *nwo.Network
		chaincode nwo.Chaincode
		process   ifrit.Process

		endorsementPluginPath string
		validationPluginPath  string
	)

	BeforeEach(func() {
		var err error
		testDir, err = ioutil.TempDir("", "pluggable-suite")
		Expect(err).NotTo(HaveOccurred())

		// Compile plugins
		endorsementPluginPath = compilePlugin("endorsement")
		validationPluginPath = compilePlugin("validation")

		// Create directories for endorsement and validation activation
		dir := filepath.Join(testDir, "endorsement")
		err = os.Mkdir(dir, 0700)
		Expect(err).NotTo(HaveOccurred())
		SetEndorsementPluginActivationFolder(dir)

		dir = filepath.Join(testDir, "validation")
		err = os.Mkdir(dir, 0700)
		Expect(err).NotTo(HaveOccurred())
		SetValidationPluginActivationFolder(dir)

		// Speed up test by reducing the number of peers we bring up
		soloConfig := nwo.BasicSolo()
		soloConfig.RemovePeer("Org1", "peer1")
		soloConfig.RemovePeer("Org2", "peer1")
		Expect(soloConfig.Peers).To(HaveLen(2))

		// docker client
		client, err = docker.NewClientFromEnv()
		Expect(err).NotTo(HaveOccurred())

		network = nwo.New(soloConfig, testDir, client, StartPort(), components)
		network.GenerateConfigTree()

		// modify config
		configurePlugins(network, endorsementPluginPath, validationPluginPath)

		// generate network config
		network.Bootstrap()

		networkRunner := network.NetworkGroupRunner()
		process = ifrit.Invoke(networkRunner)
		Eventually(process.Ready(), network.EventuallyTimeout).Should(BeClosed())

		chaincode = nwo.Chaincode{
			Name:            "mycc",
			Version:         "0.0",
			Path:            components.Build("github.com/hyperledger/fabric/integration/chaincode/module"),
			Lang:            "binary",
			PackageFile:     filepath.Join(testDir, "modulecc.tar.gz"),
			Ctor:            `{"Args":["init","a","100","b","200"]}`,
			SignaturePolicy: `OR ('Org1MSP.member','Org2MSP.member')`,
			Sequence:        "1",
			InitRequired:    true,
			Label:           "my_prebuilt_chaincode",
		}
		orderer := network.Orderer("orderer")
		network.CreateAndJoinChannel(orderer, "testchannel")
		nwo.EnableCapabilities(network, "testchannel", "Application", "V2_0", orderer, network.Peer("Org1", "peer0"), network.Peer("Org2", "peer0"))
		nwo.DeployChaincode(network, "testchannel", orderer, chaincode)
	})
