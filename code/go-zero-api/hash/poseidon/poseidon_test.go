package poseidon

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"blockchain-crypto/hash/poseidon/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPoseidonHash(t *testing.T) {
	b0 := big.NewInt(0)
	b1 := big.NewInt(1)
	b2 := big.NewInt(2)
	b3 := big.NewInt(3)
	b4 := big.NewInt(4)
	b5 := big.NewInt(5)
	b6 := big.NewInt(6)
	b7 := big.NewInt(7)
	b8 := big.NewInt(8)
	b9 := big.NewInt(9)
	b10 := big.NewInt(10)
	b11 := big.NewInt(11)
	b12 := big.NewInt(12)
	b13 := big.NewInt(13)
	b14 := big.NewInt(14)
	b15 := big.NewInt(15)
	b16 := big.NewInt(16)

	h, err := Hash([]*big.Int{b1})
	assert.Nil(t, err)
	assert.Equal(t,
		"18586133768512220936620570745912940619677854269274689475585506675881198879027",
		h.String())

	h, err = Hash([]*big.Int{b1, b2})
	assert.Nil(t, err)
	assert.Equal(t,
		"7853200120776062878684798364095072458815029376092732009249414926327459813530",
		h.String())

	h, err = Hash([]*big.Int{b1, b2, b0, b0, b0})
	assert.Nil(t, err)
	assert.Equal(t,
		"1018317224307729531995786483840663576608797660851238720571059489595066344487",
		h.String())
	h, err = Hash([]*big.Int{b1, b2, b0, b0, b0, b0})
	assert.Nil(t, err)
	assert.Equal(t,
		"15336558801450556532856248569924170992202208561737609669134139141992924267169",
		h.String())

	h, err = Hash([]*big.Int{b3, b4, b0, b0, b0})
	assert.Nil(t, err)
	assert.Equal(t,
		"5811595552068139067952687508729883632420015185677766880877743348592482390548",
		h.String())
	h, err = Hash([]*big.Int{b3, b4, b0, b0, b0, b0})
	assert.Nil(t, err)
	assert.Equal(t,
		"12263118664590987767234828103155242843640892839966517009184493198782366909018",
		h.String())

	h, err = Hash([]*big.Int{b1, b2, b3, b4, b5, b6})
	assert.Nil(t, err)
	assert.Equal(t,
		"20400040500897583745843009878988256314335038853985262692600694741116813247201",
		h.String())

	h, err = Hash([]*big.Int{b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11, b12, b13, b14})
	assert.Nil(t, err)
	assert.Equal(t,
		"8354478399926161176778659061636406690034081872658507739535256090879947077494",
		h.String())

	h, err = Hash([]*big.Int{b1, b2, b3, b4, b5, b6, b7, b8, b9, b0, b0, b0, b0, b0})
	assert.Nil(t, err)
	assert.Equal(t,
		"5540388656744764564518487011617040650780060800286365721923524861648744699539",
		h.String())

	h, err = Hash([]*big.Int{b1, b2, b3, b4, b5, b6, b7, b8, b9, b0, b0, b0, b0, b0, b0, b0})
	assert.Nil(t, err)
	assert.Equal(t,
		"11882816200654282475720830292386643970958445617880627439994635298904836126497",
		h.String())

	h, err = Hash([]*big.Int{b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11, b12, b13, b14, b15, b16})
	assert.Nil(t, err)
	assert.Equal(t,
		"9989051620750914585850546081941653841776809718687451684622678807385399211877",
		h.String())
}

func TestErrorInputs(t *testing.T) {
	b0 := big.NewInt(0)
	b1 := big.NewInt(1)
	b2 := big.NewInt(2)

	var err error

	_, err = Hash([]*big.Int{b1, b2, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0})
	require.Nil(t, err)

	_, err = Hash([]*big.Int{b1, b2, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0})
	require.NotNil(t, err)
	assert.Equal(t, "invalid inputs length 17, max 16", err.Error())

	_, err = Hash([]*big.Int{b1, b2, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0, b0})
	require.NotNil(t, err)
	assert.Equal(t, "invalid inputs length 18, max 16", err.Error())
}

func TestInputsNotInField(t *testing.T) {
	var err error

	// Very big number, should just return error and not go into endless loop
	b1 := utils.NewIntFromString("12242166908188651009877250812424843524687801523336557272219921456462821518061999999999999999999999999999999999999999999999999999999999") //nolint:lll
	_, err = Hash([]*big.Int{b1})
	require.Error(t, err, "inputs values not inside Finite Field")

	// Finite Field const Q, should return error
	b2 := utils.NewIntFromString("21888242871839275222246405745257275088548364400416034343698204186575808495617") //nolint:lll
	_, err = Hash([]*big.Int{b2})
	require.Error(t, err, "inputs values not inside Finite Field")
}

func TestHashBytes(t *testing.T) {
	type testVector struct {
		bytes        string
		expectedHash string
	}
	//nolint:lll
	var testVectors = []testVector{
		{
			bytes:        "dead",
			expectedHash: "0244ec1a137a24c92404de9f9c39907be151026a4eb7f9cfea60a5740e8a73b7",
		},
		{
			bytes:        "ce0e8e502600b14bb0a2c9689f7d93d10e9f5451f18f030ec3bb6c5001ea60",
			expectedHash: "03ace829d65dae7903fc9b22e904750ca0509b9a187cbeb004a4270e8322a2f4",
		},
		{
			bytes:        "34665289b71a2cb8bf4c289ae6d17d845457c48bfc18623ca39e141b2e40c5d3",
			expectedHash: "144aa8e107a9fd8f8106016a8ff0fd1c265cdf31a40a8ef9ed44dfb4592cbbff",
		},
		{
			bytes:        "3211bcce6a7d8132020223eef1a03385ba6bd4966b295c2e2211a8d6d9e389fe6bf08f21497774456be2e47fdb6740aa571338c71c38c0a6d7f703007569e64031633ec7c8ef2ba25ad6a248403deb697457fae8a4a7f4525d73a3d4cd93334a894efbb20d0a6391df0aae46bc32005834ed084aeb08887e08eb67cde004fea6f8036b061fa8cb7246af2458a4cef79c648b13ef8ac50d9a8863be1c58a7a9a5940006022611ca35508b993656cf3fd0175579c6983414701134cc0becc51364289d4775b71b67f269a16fe653a00ab75885924777feaa990cce9c561802581b9092e9be2a0d03fd86361e427b94d8600a7edc67c263b35a0be6837e750175b50314c7d4642534b3233c963e397f63e6d7187b114eef1346412de83993cb79bc80e9a921fa59ccccda30e57025ccaa0830e1eb1ea5c87ca6fc887aedabdab1bb4cf6022440960b0e03f5de85137d48392873851d13f8035b67e6a5f5c5bac7598fe2f91673f3875b40faad43357862b76e9c6062b3342f199bec165e3093b8c25e21ac626d718e8aaa0d8aacc034a2da4a6ff3de36891ddd30b22abedf0f72b493e9f16aaa65fddeff83612b1d07989e1d6d1ba7600123645c5920f55678cf518d8f58d73d6227e710bcf6dfcfe309c5d67e4f51fbb18aa3922c07c35e5fefa66c0c57553d5ab9e323591031ecfb0b84",
			expectedHash: "25c9473a28eaf3d42fa89c218ac91da88311cd762272ddedc79645696face6d8",
		},
		{
			bytes:        "e7190c27f2adfb643bdbaa686b8372df9e8132d079640d43f218bfe3b8fafacfe0d7855d77cc195e05d855d2d828e627be305b34f4de52894d3672515f2e1dcfd4e6909a5406d5dbdff31e38d2400299dc6a5f2509052c76393d57a786afc51c18d63c9bd2edd70b8d82f867c5269556a636d1c8c7fb45884adf264c9ae64731dfac21efb29dcddb5bef66149f7af55f0263dc16f7cd8f58f3eb1f1b5246a76ae64ccf7731df0e17963efa4b786d3365a2f2adfd87767d3bdfa104c443c2c0eb79d408cf0469b592f2863988bee8b42e9955255c3d632edf1a3de2a305f573112575f4958192f1b433089fd928dc8f38263f43be6d7cb83acbe4ac2bc0d44f219edbfcafaa29fabda4c8b41add2525fa38982649fc22ef1221273441e65ce62ea200ed951a1619ee6c053793096788c09406b2f9bd09d579dc1fef5f44ba91460f93aaf278bf3f4d25536c1ecf3af64f83a7a04ec049e54ffa007721cfc1a336089824bff3a23f39421234ba1a5f6113eae0bbfbf6a9295d5d473838a6dae3e34620bf365ff588a1ebce6c5bf9c46038a5f323d23ba7ec5d9afc48127612eec3620fff7472dc342f4e56d5e36fd910a66ec8d95ead3f06eb47f612063ea4d64b90bcf5f199684e99f98732029478d99505ca73b86e0cf4b51c63d9cdb24fb4f54908e4cb98aec0af7587b2a4b9477061f2",
			expectedHash: "051a6d9a46fe012bc06cc958a1337cb41a4d7b91f9400cea458f37af18be222e",
		},
		{
			bytes:        "8246af280f1863b4eabf37a548ce764a296fd81be5b23c4fc04ee3540c81a765c795c8beedc4a6a1d67a53fcc9525bc2e2d31369fc64e0f547a4de44bd312af1288604fd183dbb52bbdd445ce870c24b829007162a438eb22fbfb08939f4b314f86f264d17126cfd1cd50028eec5aabd9bb1f5c759938b02f0d9d1cbccc4655735b65cf80cad3ce3fab37fac5833652d68d633a3d12ce024e9d16d1fa8a0c5ba8072286e855594471255dffe96a87568501813bac166a92f356ab38032097e5a68406bc22faee6db58f5ae36f24e877ea72a5dc6978c4a5c7f671e635ec4430da1fd2e9e9587b2a8b64b841108f3f5c7af3be9fef9e940478b021352055320d55bb2bf292a1e796154c5e530284a16dda5aeb29baca584f76487cf20f5bb409cfb6247a6918ebf8df1c854551a2184cd06df5706b2fb4d92a70cece8eb4bcee91934f09c0310efcbff2dbc9aa5b6ea818d96f471f9025cb46e4acd98a4ecc6dd2d647282f05a2c586ffe4b94ac12f36bd65cf6f3903d0228f92df340afb425acc5df8433407de697a44a7780506896799c5dca139ed9498880c5b17739859c30b4caa7945984d3d7d4c8450ad1c0f37c50f280fce45f791ffe12b199eab24eca47223efcdc86d764b57fb3afcdc4b01b6f35a3773f0331912bda3ba36d705a8ea506c99d573255233f13eba6d88280d4bf26d24419fc9b93139b74880d407cb796c2d6d1385b6a456dd7ab5755a8465972ae1d2e34d50a302bb1617cd75251f83af0f2be482d29b78b9b4320accf1c45407bc1b2c6dbbd26a8d8811ca95bd88e59bed19a163ab88a5d9b3e286eb3b4cf7ea388d32d613e0f16331570b93fe79ebc221e8fa8eb22e08b205237435e5395198bef4d264953b3fbd72c761603f1b343e62363369dd3f1c382487655fdeb6aba314f500466fea3293ef6feef458bfaec77f6f1f3ce3525b7fe2df433b07330d179eb427739782c95c8767ab113444208a30a4e78e6e18869972a412ef28f3925bcb857e0716e66814dc31abc37bf20219eb9c60f35ef4e1f10b73889c9094867d813a158d0acb38239da0e3f6d1d8865fe49099dfdb6cb7c55160ddb81d0d0a431eec0ecc19f878cf92f2a9a58d951e5b4a8e3b8e87756577b157aec0b3911fa38814fc752c9377b98bc4477172ba2a9823f33a00d50fde16b148c9815f70e7057e181c757ecded9b0e31d35d2c7707dde7d855e6cf2b5f5496229",
			expectedHash: "189569b8bb6878bc0eac1136913cd74eb5fa744bd54c7fc473ffe333e196bfb2",
		},
		{
			bytes:        "24bccf22f7178476bc30c4cb690bb1df362e258bb07a0e862910568907815b5e276b7e3df3e94d45865f37ba13af512f42afe02a5bffc775e85f3483bf26a60779985d3713030544c9881a54e81549fdf0efa37529b3fc4416fead3efbaa921b88d182ca2cae3f0552a22533e1be3663708e4ef91ad82a54aa3b3f3bb6a24a2a296f320086b6e9a10dec3e5f67b8cc358ca5086d8f3a84fb0ca90167964a9cea26aa31fab862d86263221f67e52ee5fc4a29a0b6e04472f718e76b6a19f6493d540da2513a5547f9649a4f77dc99e0bb0b1e680da69fa2dd9ec2b3b11e375bc993d300adfcd7fc4fe875aa4d29780015fb77c7bd135aa49e62917043147a3188ece587f88da06fe698793d1e3b1d7cedd93b12664f53defeebccdb816f91fad2fff1214c97e61d1ac6178f278273d656a209a4eb1f532fc34ce808932ade305c27a21bed0bfe4595364f6e6462732f75fc8566bb8d1d6fec34228f510cbb1f72293c5d3cd9cb38b1f77db7d67fef1032f392b156c3a1c8368307bcc1843b1a7661630417289c61d44457ad6455c84d73db4517e771ea8a2041113e9910c9dc64bfbac9e7f6563014c65dce227f4f11b8071385dab30bc5c4e3553c4919a1f14302b491106844ce5b7b5874d879fed8e87ff3edf3ae04f2b77ad4a18051317cac773c3800bdfacf59a027a876e471b8c35d348a6690d6241d47119fb0c44448f03e2030c6999a48f3010e6d85503564d3af954671d0921fcd64d37b937a40d9070ca1b9e41f4cfc1e29b9946fc103418beab9fad91201cf00c13a5900ea2aa4d174fc7bbbbef1cd34b8b4f5853ebed9457fed88645775c8b3f466a474746338bd2a37263e518a538db50e74c9f4ce996e3e6ccd5153d3f2637f51aeb27018f0faf9ee01766c14ba860a689502459878e7383a9e314921237dcd968ea03083c8e66856789dd5e9a8d0b1dcdca37c9262391bf5943f5fcd8fc24637e23b91ccefe485dc19c4ea645a14c5f586be606488eea01307c886b7c8f976071c58c2e8f5c4e494a4f437d5696271fdd7ff0d85da03f0d151e5061676887641a2d178cac44af38e0588bfacfb71c797761d72fdd2a25ee1a181548f95d410cdbab2968fee3b35aee8cdb9fefc344f8cc554384c2aaf25b8277ab965d2fc27651361c7a23608277a2fb22230a15cfe69860d219603dd37f3d819f5877aa732ea67d36f6457614246f8ff11a3ef18d640578f2f95887a4dcc1152e02e51aff0b1a807c6973bb47eb315d922a8fbb00eb926b4c46827f9c48190fbb94d4795e3dfbe1b878e15b5d8dddace7f82451de45cf564a85beadf1768ee25bba3e1f7f41ea85f07534e67e2",
			expectedHash: "2a61df8c7495a7cb0f575c457c21c072bc69ebd1393aebd903c9a882df1b1d7b",
		},
		{
			bytes:        "ab8ffd3224ad21c4009f6e67a9047209f338f88f4f29bcffcb769695d0963b859c591f85b570858c2eb16262a3c07ca323ed0ae762c6772e4f3494036c82a71f1b5569be153d1b069ceecd5fd88f2055cda4d9a651b70c7f736a4fb5a8eb383db80530e608ce01fe52956ba5ebe2801ae60309c4d84021839544933559178ca84b9cbb53cff9a831d6098d75ed580352f70577479c266443686923596280a36e8b811b858f69d79966d2548ab2d830fd10010b968889abbecb865af6d9f6ae6e95353900c2c260a72f3ae32660c14491ba7a97f9896f63f7f29d021212a4c18148a386ef60b6b16d834645e2fe6c9279fabc0056168aab2c023ea3443bc4ea62185f03730121e2df8cd015cadae0b1dd95f723f738af08c9f8d7d47f35d64e580ad7d6525a5b9648543fd752d5c27e3c5eefde55a3d867d98107da151399d6b8718cf662e07562025447b37ebf3854edee2fb74c51273d775ad4a4ece04ec840ffb3af5e184927ee644ce5a3c872ca354bef7187797062206e8b1d97b599401aca9b401d72e2611ccfdd73c594a7c4aa8cabd3d431212ccc1512efd0cca758d9650b53fc3cda49c74f6a048b947a3f48e92c2702a2f1d585817820a9ecefe3fc90332a718851c96a761d61395d84bf7866862ea57c8d49a4e417afdae540b8f607eac7bc4450a3ec30e28c8889bd50a036c03e633735b793dd6030880661e0a1756937ff9556a52de804ef5fc381b29b73f8bfbd13a8b2155ad919bca8ee1f5ce6d10a608154e6f79faca4d5eda068bf71bf31f944ca5c98baca4a5d7a141d9001dcb18037e30667eed18e39a1eb2ad0124c0dc6b4d3afe421baea72d2c15a658c5651c150b67e79c599839ca00af436833f65db98635959157f70446ddf578ec1a71c485f0bfbb071b23226a1ca34da34cf66c09d969dd01adc65efdf9222d075e7873053e029850bab855cdc4f15180a5571d7579c8bc5f470fa4b2a9ed154c505a513e5f0867cd302d8b6e649237079b7ee0640fe2b5ac8de694da87221a7177dc316c2ea47c1edb6a7699f3651a99f43ce2a885146d006bcb82864508223213d930e963b1305b7255807bc5f93498d832f5ed224c385253aed3b8354c511c3dafea13d3cd5e00523f00a7bf05c2a65fa23c0ac915d7098254ddc821a6f3df7b4d8aa79b4da56b4617382ed11dc8b0561fb037a10c5f9917a68ea5d1d8e7b0fec11628b8ce493b8ff7426cd6c43edc58e80b98b381be0399a51a146041f52db2eb1c4fde3cca60ede53050c9bd72e67ca41155b1ac1508ab4055b4053064af44b8e31c162d447935ef5d821ea108536733d949363e28980d78b0476a6092968f7",
			expectedHash: "2b07135db61f6943328ef48524aefabf74a44246b543d5bd90d250fafc8b0f78",
		},
		{
			bytes:        "60806040523480156200001157600080fd5b5060405162002103380380620021038339810160408190526200003491620002e1565b6200003f3362000125565b60005b620000506001602062000329565b811015620000eb57602181602081106200006e576200006e62000343565b01546021826020811062000086576200008662000343565b0154604080516020810193909352820152606001604051602081830303815290604052805190602001206021826001620000c1919062000359565b60208110620000d457620000d462000343565b015580620000e28162000374565b91505062000042565b50604680546001600160a01b0319166001600160a01b0383161790556200011162000175565b6043556200011e62000250565b50620003b5565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6041546000908190815b602081101562000248578160011660011415620001e05760018160208110620001ac57620001ac62000343565b0154604080516020810192909252810184905260600160405160208183030381529060405280519060200120925062000224565b8260218260208110620001f757620001f762000343565b01546040805160208101939093528201526060016040516020818303038152906040528051906020012092505b6200023160028362000392565b9150806200023f8162000374565b9150506200017f565b509092915050565b60458054906000620002628362000374565b909155505060435460425460408051602081019390935282015260600160408051808303601f19018152828252805160209182012060455460009081526044835283902055604354604254908452908301527f61014378f82a0d809aefaf87a8ac9505b89c321808287a6e7810f29304c1fce3910160405180910390a1565b600060208284031215620002f457600080fd5b81516001600160a01b03811681146200030c57600080fd5b9392505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156200033e576200033e62000313565b500390565b634e487b7160e01b600052603260045260246000fd5b600082198211156200036f576200036f62000313565b500190565b60006000198214156200038b576200038b62000313565b5060010190565b600082620003b057634e487b7160e01b600052601260045260246000fd5b500490565b611d3e80620003c56000396000f3fe60806040526004361061010e5760003560e01c806355f6bc57116100a55780638da5cb5b11610074578063a71d644411610059578063a71d6444146102fb578063ed6be5c91461032b578063f2fde38b1461035557600080fd5b80638da5cb5b146102b0578063a5392cf6146102db57600080fd5b806355f6bc57146101f95780635ec6a8df14610219578063715018a61461026b5780637d8f04691461028057600080fd5b8063319cf735116100e1578063319cf7351461017d5780633381fe90146101935780633ae05047146101c05780633ed691ef146101d557600080fd5b806301fd904414610113578063029f27931461013c5780630e21fbd7146101525780632dfdf0b514610167575b600080fd5b34801561011f57600080fd5b5061012960425481565b6040519081526020015b60405180910390f35b34801561014857600080fd5b5061012960455481565b61016561016036600461170e565b610375565b005b34801561017357600080fd5b5061012960415481565b34801561018957600080fd5b5061012960435481565b34801561019f57600080fd5b506101296101ae36600461175f565b60446020526000908152604090205481565b3480156101cc57600080fd5b50610129610574565b3480156101e157600080fd5b50604554600090815260446020526040902054610129565b34801561020557600080fd5b5061016561021436600461175f565b610641565b34801561022557600080fd5b506046546102469073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610133565b34801561027757600080fd5b506101656106f8565b34801561028c57600080fd5b506102a061029b36600461186d565b610785565b6040519015158152602001610133565b3480156102bc57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610246565b3480156102e757600080fd5b506101656102f6366004611915565b610931565b34801561030757600080fd5b506102a061031636600461175f565b60476020526000908152604090205460ff1681565b34801561033757600080fd5b50610340600081565b60405163ffffffff9091168152602001610133565b34801561036157600080fd5b506101656103703660046119d2565b610e4d565b73ffffffffffffffffffffffffffffffffffffffff841661042957823414610424576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4272696467653a3a6465706f7369743a20414d4f554e545f444f45535f4e4f5460448201527f5f4d415443485f4d53475f56414c55450000000000000000000000000000000060648201526084015b60405180910390fd5b61044b565b61044b73ffffffffffffffffffffffffffffffffffffffff8516333086610f7a565b63ffffffff82166104de576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f4272696467653a3a6465706f7369743a2044455354494e4154494f4e5f43414e60448201527f545f42455f4d41494e4e45540000000000000000000000000000000000000000606482015260840161041b565b6041546040805173ffffffffffffffffffffffffffffffffffffffff87811682526020820187905263ffffffff8681168385015290851660608301529092166080830152517f0a37f8bae6de7e960aeedce45875d5a75681918316c4bd81f4691152910f8e329181900360a00190a161055b848460008585611056565b610563610574565b60435561056e61125e565b50505050565b6041546000908190815b60208110156106395781600116600114156105d957600181602081106105a6576105a66119ef565b0154604080516020810192909252810184905260600160405160208183030381529060405280519060200120925061061a565b82602182602081106105ed576105ed6119ef565b01546040805160208101939093528201526060016040516020818303038152906040528051906020012092505b610625600283611a4d565b91508061063181611a88565b91505061057e565b509092915050565b60465473ffffffffffffffffffffffffffffffffffffffff1633146106e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4272696467653a3a757064617465526f6c6c757045786974526f6f743a204f4e60448201527f4c595f524f4c4c55500000000000000000000000000000000000000000000000606482015260840161041b565b60428190556106f561125e565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610779576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161041b565b610783600061130b565b565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e088811b821660208401527fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060608c811b82166024860152603885018c90529189901b909216605884015286901b16605c8201526000908190607001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120905067ffffffffffffffff841660005b60208110156109205781600116600114156108b357868181518110610873576108736119ef565b602002602001015183604051602001610896929190918252602082015260400190565b604051602081830303815290604052805190602001209250610901565b828782815181106108c6576108c66119ef565b60200260200101516040516020016108e8929190918252602082015260400190565b6040516020818303038152906040528051906020012092505b61090c600283611a4d565b91508061091881611a88565b91505061084c565b505090911498975050505050505050565b67ffffffffffffffff841660009081526047602052604090205460ff16156109db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4272696467653a3a77697468647261773a20414c52454144595f434c41494d4560448201527f445f574954484452415700000000000000000000000000000000000000000000606482015260840161041b565b63ffffffff871615610a6f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4272696467653a3a77697468647261773a2044455354494e4154494f4e5f4e4560448201527f54574f524b5f4e4f545f4d41494e4e4554000000000000000000000000000000606482015260840161041b565b63ffffffff881615610b03576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f4272696467653a3a77697468647261773a204f524947494e5f4e4554574f524b60448201527f5f4e4f545f4d41494e4e45540000000000000000000000000000000000000000606482015260840161041b565b6000838152604460209081526040918290205482519182018590529181018390526060016040516020818303038152906040528051906020012014610bca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4272696467653a3a77697468647261773a20474c4f42414c5f455849545f524f60448201527f4f545f444f45535f4e4f545f4d41544348000000000000000000000000000000606482015260840161041b565b610bda8a8a8a8a8a8a8a88610785565b610c40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4272696467653a3a77697468647261773a20534d545f494e56414c4944000000604482015260640161041b565b67ffffffffffffffff8416600090815260476020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905573ffffffffffffffffffffffffffffffffffffffff8a16610dad576040805160008082526020820190925273ffffffffffffffffffffffffffffffffffffffff8816908b90604051610cd49190611aed565b60006040518083038185875af1925050503d8060008114610d11576040519150601f19603f3d011682016040523d82523d6000602084013e610d16565b606091505b5050905080610da7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4272696467653a3a77697468647261773a204554485f5452414e534645525f4660448201527f41494c4544000000000000000000000000000000000000000000000000000000606482015260840161041b565b50610dce565b610dce73ffffffffffffffffffffffffffffffffffffffff8b16878b611380565b6040805167ffffffffffffffff8616815263ffffffff8a16602082015273ffffffffffffffffffffffffffffffffffffffff8c811682840152606082018c90528816608082015290517f8932892d010aea7e4fdefb3764910523c321e06bb52577dc2439501196bf72559181900360a00190a150505050505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610ece576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161041b565b73ffffffffffffffffffffffffffffffffffffffff8116610f71576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161041b565b6106f58161130b565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261056e9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526113db565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085811b821660208401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606089811b82166024860152603885018990529186901b909216605884015283901b16605c8201526000906070016040516020818303038152906040528051906020012090506001602060026110fd9190611c2b565b6111079190611c37565b60415410611197576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4465706f736974436f6e74726163743a5f6465706f7369743a204d45524b4c4560448201527f5f545245455f46554c4c00000000000000000000000000000000000000000000606482015260840161041b565b6001604160008282546111aa9190611c4e565b909155505060415460005b602081101561124b5781600116600114156111e95782600182602081106111de576111de6119ef565b015550611257915050565b600181602081106111fc576111fc6119ef565b015460408051602081019290925281018490526060016040516020818303038152906040528051906020012092506002826112379190611a4d565b91508061124381611a88565b9150506111b5565b50611254611c66565b50505b5050505050565b6045805490600061126e83611a88565b9091555050604354604254604080516020810193909352820152606001604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152828252805160209182012060455460009081526044835283902055604354604254908452908301527f61014378f82a0d809aefaf87a8ac9505b89c321808287a6e7810f29304c1fce3910160405180910390a1565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526113d69084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401610fd4565b505050565b600061143d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166114e79092919063ffffffff16565b8051909150156113d6578080602001905181019061145b9190611c95565b6113d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161041b565b60606114f68484600085611500565b90505b9392505050565b606082471015611592576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161041b565b843b6115fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161041b565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516116239190611aed565b60006040518083038185875af1925050503d8060008114611660576040519150601f19603f3d011682016040523d82523d6000602084013e611665565b606091505b5091509150611675828286611680565b979650505050505050565b6060831561168f5750816114f9565b82511561169f5782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041b9190611cb7565b73ffffffffffffffffffffffffffffffffffffffff811681146106f557600080fd5b803563ffffffff8116811461170957600080fd5b919050565b6000806000806080858703121561172457600080fd5b843561172f816116d3565b935060208501359250611744604086016116f5565b91506060850135611754816116d3565b939692955090935050565b60006020828403121561177157600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126117b857600080fd5b8135602067ffffffffffffffff808311156117d5576117d5611778565b8260051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f8301168101818110848211171561181857611818611778565b60405293845285810183019383810192508785111561183657600080fd5b83870191505b848210156116755781358352918301919083019061183c565b803567ffffffffffffffff8116811461170957600080fd5b600080600080600080600080610100898b03121561188a57600080fd5b8835611895816116d3565b9750602089013596506118aa60408a016116f5565b95506118b860608a016116f5565b945060808901356118c8816116d3565b935060a089013567ffffffffffffffff8111156118e457600080fd5b6118f08b828c016117a7565b9350506118ff60c08a01611855565b915060e089013590509295985092959890939650565b6000806000806000806000806000806101408b8d03121561193557600080fd5b8a35611940816116d3565b995060208b0135985061195560408c016116f5565b975061196360608c016116f5565b965060808b0135611973816116d3565b955060a08b013567ffffffffffffffff81111561198f57600080fd5b61199b8d828e016117a7565b9550506119aa60c08c01611855565b935060e08b013592506101008b013591506101208b013590509295989b9194979a5092959850565b6000602082840312156119e457600080fd5b81356114f9816116d3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082611a83577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611aba57611aba611a1e565b5060010190565b60005b83811015611adc578181015183820152602001611ac4565b8381111561056e5750506000910152565b60008251611aff818460208701611ac1565b9190910192915050565b600181815b80851115611b6257817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115611b4857611b48611a1e565b80851615611b5557918102915b93841c9390800290611b0e565b509250929050565b600082611b7957506001611c25565b81611b8657506000611c25565b8160018114611b9c5760028114611ba657611bc2565b6001915050611c25565b60ff841115611bb757611bb7611a1e565b50506001821b611c25565b5060208310610133831016604e8410600b8410161715611be5575081810a611c25565b611bef8383611b09565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115611c2157611c21611a1e565b0290505b92915050565b60006114f98383611b6a565b600082821015611c4957611c49611a1e565b500390565b60008219821115611c6157611c61611a1e565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600060208284031215611ca757600080fd5b815180151581146114f957600080fd5b6020815260008251806020840152611cd6816040850160208701611ac1565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fea264697066735822122086e50005a89e4c2272c1cba11ff802c1213e2b08c8a20c742826166345c5ee3264736f6c63430008090033",
			expectedHash: "13a3be23e0f86283cb0ba23739fb1020aa7c85ae498f8d543e37cadcd5ec408b",
		},
	}
	for i, vector := range testVectors {
		t.Run(fmt.Sprintf("test vector %d", i), func(t *testing.T) {
			inputBytes, err := hex.DecodeString(vector.bytes)
			require.NoError(t, err)
			res, err := HashBytes(inputBytes)
			require.NoError(t, err)
			require.NotEmpty(t, res)
			assert.Equal(t, vector.expectedHash, hex.EncodeToString(res.Bytes()))
		})
	}
}

func BenchmarkPoseidonHash6Inputs(b *testing.B) {
	b0 := big.NewInt(0)
	b1 := utils.NewIntFromString("12242166908188651009877250812424843524687801523336557272219921456462821518061") //nolint:lll
	b2 := utils.NewIntFromString("12242166908188651009877250812424843524687801523336557272219921456462821518061") //nolint:lll

	bigArray6 := []*big.Int{b1, b2, b0, b0, b0, b0}

	for i := 0; i < b.N; i++ {
		Hash(bigArray6) //nolint:errcheck,gosec
	}
}

func BenchmarkPoseidonHash8Inputs(b *testing.B) {
	bigArray8 := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
		big.NewInt(4),
		big.NewInt(5),
		big.NewInt(6),
		big.NewInt(7),
		big.NewInt(8),
	}

	for i := 0; i < b.N; i++ {
		Hash(bigArray8) //nolint:errcheck,gosec
	}
}

func BenchmarkPoseidonHash16Inputs(b *testing.B) {
	bigArray16 := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
		big.NewInt(4),
		big.NewInt(5),
		big.NewInt(6),
		big.NewInt(7),
		big.NewInt(8),
		big.NewInt(9),
		big.NewInt(10),
		big.NewInt(11),
		big.NewInt(12),
		big.NewInt(13),
		big.NewInt(14),
		big.NewInt(15),
		big.NewInt(16),
	}

	for i := 0; i < b.N; i++ {
		Hash(bigArray16) //nolint:errcheck,gosec
	}
}
