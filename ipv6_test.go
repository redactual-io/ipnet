package ipnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_v6IP_addr(t *testing.T) {
	type fields struct {
		address address
	}
	tests := []struct {
		name   string
		fields fields
		want   address
	}{
		{
			"Test1",
			fields{
				address{0xff02, 0x1},
			},
			address{0xff02, 0x1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := v6IP{
				address: tt.fields.address,
			}
			assert.Equalf(t, tt.want, i.addr(), "addr()")
		})
	}
}

func Test_v6IP_Bytes(t *testing.T) {
	type fields struct {
		address address
	}
	tests := []struct {
		name   string
		fields fields
		want   [8]uint64
	}{
		{
			"Test1",
			fields{
				address{0x1234567890abcdef, 0x1234567890abcdef},
			},
			[8]uint64{
				0x1234,
				0x5678,
				0x90ab,
				0xcdef,
				0x1234,
				0x5678,
				0x90ab,
				0xcdef,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := v6IP{
				address: tt.fields.address,
			}
			assert.Equalf(t, tt.want, i.bytes(), "bytes()")
		})
	}
}

func Test_v6IP_String(t *testing.T) {
	type fields struct {
		address address
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Test1",
			fields{
				address{0x1234567890abcdef, 0x1234567890abcdef},
			},
			"1234:5678:90ab:cdef:1234:5678:90ab:cdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := v6IP{
				address: tt.fields.address,
			}
			assert.Equalf(t, tt.want, i.String(), "String()")
		})
	}
}

func Test_v6IP_Version(t *testing.T) {
	type fields struct {
		address address
	}
	tests := []struct {
		name   string
		fields fields
		want   Ver
	}{
		{
			"Test1",
			fields{
				address{0xff02, 0x1},
			},
			v6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := v6IP{
				address: tt.fields.address,
			}
			assert.Equalf(t, tt.want, i.Version(), "Version()")
		})
	}
}

func Test_splitV6(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		args   args
		want   []string
		wantOk bool
	}{
		{
			"ShorthandFail",
			args{
				"ff02::1",
			},
			[]string{},
			false,
		},
		{
			"TooManyHextets",
			args{
				"ff02:0:0:0:0:0:0:1:1",
			},
			[]string{},
			false,
		},
		{
			"HextetTooLong",
			args{
				"ff02:0:0:0:0:0:0:11111",
			},
			[]string{},
			false,
		},
		{
			"Empty",
			args{
				"",
			},
			[]string{},
			false,
		},
		{
			"EmptyEndBlock",
			args{
				"2006:1:",
			},
			[]string{},
			false,
		},
		{
			"Valid1",
			args{
				"1234:5678:90ab:cdef:1234:5678:90ab:cdef",
			},
			[]string{"1234", "5678", "90ab", "cdef", "1234", "5678", "90ab", "cdef"},
			true,
		},
		{
			"Valid2",
			args{
				"2602:8084:3:6580:68f0:fee6:7936:2667",
			},
			[]string{"2602", "8084", "3", "6580", "68f0", "fee6", "7936", "2667"},
			true,
		},
		{
			"Valid3",
			args{
				"1234:5678:90ab:cdef",
			},
			[]string{"1234", "5678", "90ab", "cdef"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitV6(tt.args.s)
			assert.Equalf(t, tt.want, got, "splitV6(%v)", tt.args.s)
			assert.Equalf(t, tt.wantOk, got1, "splitV6(%v)", tt.args.s)
		})
	}
}

func Test_splitShortV6(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		args   args
		want   []string
		wantOk bool
	}{
		{
			"Valid1",
			args{
				"ff02::1",
			},
			[]string{"ff02", "0", "0", "0", "0", "0", "0", "1"},
			true,
		},
		{
			"Valid2",
			args{
				"2001::4675:b261",
			},
			[]string{"2001", "0", "0", "0", "0", "0", "4675", "b261"},
			true,
		},
		{
			"Valid3",
			args{
				"2400:cb00:2049:1::629f:16e6",
			},
			[]string{"2400", "cb00", "2049", "1", "0", "0", "629f", "16e6"},
			true,
		},
		{
			"Valid4",
			args{
				"2606:6000::",
			},
			[]string{"2606", "6000", "0", "0", "0", "0", "0", "0"},
			true,
		},
		{
			"Valid5",
			args{
				"2605:6440:3008:9000::64cc",
			},
			[]string{"2605", "6440", "3008", "9000", "0", "0", "0", "64cc"},
			true,
		},
		{
			"Invalid1",
			args{
				"26051:6440:3008:9000::64cc",
			},
			[]string{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitShortV6(tt.args.s)
			assert.Equalf(t, tt.want, got, "splitShortV6(%v)", tt.args.s)
			assert.Equalf(t, tt.wantOk, got1, "splitShortV6(%v)", tt.args.s)
		})
	}
}

func Test_stretchV6(t *testing.T) {
	type args struct {
		o [][]string
	}
	tests := []struct {
		name   string
		args   args
		want   []string
		wantOk bool
	}{
		{
			"Test1",
			args{
				[][]string{{"ff02"}, {"1"}},
			},
			[]string{"ff02", "0", "0", "0", "0", "0", "0", "1"},
			true,
		},
		{
			"Test2",
			args{
				[][]string{{"ff02", "0", "0", "0", "0", "0", "0", "1"}, {"1"}},
			},
			[]string{},
			false,
		},
		{
			"Test3",
			args{
				[][]string{{"ff02", "0", "0", "0", "0", "0"}, {"1"}},
			},
			[]string{"ff02", "0", "0", "0", "0", "0", "0", "1"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := stretchV6(tt.args.o)
			assert.Equalf(t, tt.want, got, "stretchV6(%v)", tt.args.o)
			assert.Equalf(t, tt.wantOk, gotOk, "stretchV6(%v)", tt.args.o)
		})
	}
}

func Test_parseV6(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		args   args
		want   [8]uint64
		wantOk bool
	}{
		{
			"Test1",
			args{
				"ff02::1",
			},
			[8]uint64{0xff02, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
			true,
		},
		{
			"Test2",
			args{
				"ff02:0:0:0:0:0:0:1",
			},
			[8]uint64{0xff02, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
			true,
		},
		{
			"TooManyHextets",
			args{
				"ff02:0:0:0:0:0:0:0:1",
			},
			[8]uint64{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseV6(tt.args.s)
			assert.Equalf(t, tt.want, got, "parseV6(%v)", tt.args.s)
			assert.Equalf(t, tt.wantOk, got1, "parseV6(%v)", tt.args.s)
		})
	}
}

func Test_atoUint64(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name   string
		args   args
		want   [8]uint64
		wantOk bool
	}{
		{
			"Test1",
			args{
				[]string{"ff02", "0", "0", "0", "0", "0", "0", "1"},
			},
			[8]uint64{0xff02, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
			true,
		},
		{
			"Test2",
			args{
				[]string{"ff02", "z", "0", "0", "0", "0", "0", "1"},
			},
			[8]uint64{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			false,
		},
		{
			"Test3",
			args{
				[]string{"ff02", "0", "0", "0", "1"},
			},
			[8]uint64{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := a16toUint64(tt.args.s)
			assert.Equalf(t, tt.want, got, "a16toUint64(%v)", tt.args.s)
			assert.Equalf(t, tt.wantOk, got1, "a16toUint64(%v)", tt.args.s)
		})
	}
}

func Test_parseIPv6(t *testing.T) {
	type args struct {
		addrStr string
	}
	tests := []struct {
		name    string
		args    args
		want    v6IP
		wantErr bool
	}{
		{
			"Test1",
			args{
				"ff02::1",
			},
			v6IP{address: address{0xff02000000000000, 0x1}},
			false,
		},
		{
			"Test2",
			args{
				"f",
			},
			v6IP{},
			true,
		},
		{
			"Test3",
			args{
				"",
			},
			v6IP{},
			true,
		},
		{
			"Test4",
			args{
				"ff022::",
			},
			v6IP{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseIPv6(tt.args.addrStr)
			if (err != nil) != tt.wantErr {
				t.Error(err)
			}
			assert.Equalf(t, tt.want, got, "parseIPv6(%v)", tt.args.addrStr)
		})
	}
}

func Test_parseIPv6Sanity(t *testing.T) {
	for testAddr, expected := range testv6addr {
		got, err := parseIPv6(testAddr)
		assert.Nil(t, err)
		assert.Equal(t, expected, got.addr())
	}
}

var testv6addr = map[string]address{
	"2602:8084:3:6580:68f0:fee6:7936:2667":    {0x2602808400036580, 0x68f0fee679362667},
	"2001:4176:2:1359:65:236:44:166":          {0x2001417600021359, 0x65023600440166},
	"2600:1f16:2b4:6200:9697:f425:63c4:7214":  {0x26001f1602b46200, 0x9697f42563c47214},
	"2600:9000:218d:b600:1f:62b2:85c0:9361":   {0x26009000218db600, 0x1f62b285c09361},
	"2601:151:8300:d070:81d8:e631:4078:db7f":  {0x260101518300d070, 0x81d8e6314078db7f},
	"2601:197:4400:3b26:5412:65eb:8fe6:3210":  {0x2601019744003b26, 0x541265eb8fe63210},
	"2602:26f0:41:69c::3962":                  {0x260226f00041069c, 0x3962},
	"2602:4780:6:679::1e57:9942:d":            {0x2602478000060679, 0x1e579942000d},
	"2601:243:c002:dd6d:6c45:ce76:6cfb:bf4c":  {0x26010243c002dd6d, 0x6c45ce766cfbbf4c},
	"2606:4700:50::629f:266c":                 {0x2606470000500000, 0x629f266c},
	"2601:e06:2bc:f010:c1e1:6294:6d7c:7780":   {0x26010e0602bcf010, 0xc1e162946d7c7780},
	"2003:c1:73f1:6876:fd63:4851:3d05:5db8":   {0x200300c173f16876, 0xfd6348513d055db8},
	"2003:c6:5f0e:1b29:bc87:30d6:6607:bd42":   {0x200300c65f0e1b29, 0xbc8730d66607bd42},
	"2409:4041:2696:f29e:d43:b66d:56b4:4486":  {0x240940412696f29e, 0xd43b66d56b44486},
	"2600:1702:2621:110:2527:1c63:6123:b0b7":  {0x2600170226210110, 0x25271c636123b0b7},
	"2600:1f14:35:3002:92ed:24e2:698d:dfb2":   {0x26001f1400353002, 0x92ed24e2698ddfb2},
	"2601:681:600:e3e0:1466:c91b:83c2:3ecd":   {0x260106810600e3e0, 0x1466c91b83c23ecd},
	"2600:9000:2171:ee00:15:c08:1bc0:9361":    {0x260090002171ee00, 0x150c081bc09361},
	"2800:810:42b:df6:6555:74c1:f30b:c1d4":    {0x28000810042b0df6, 0x655574c1f30bc1d4},
	"2601:4f8:212:3122:1000:2000::2":          {0x260104f802123122, 0x1000200000000002},
	"2607:5300:263:755b::":                    {0x260753000263755b, 0x0},
	"2001:8d8:fe:53::d960:501f:100":           {0x200108d800fe0053, 0xd960501f0100},
	"2605:e000:1303:4166:f927:348e:410f:6417": {0x2605e00013034166, 0xf927348e410f6417},
	"2606:4700:10::6c43:75e":                  {0x2606470000100000, 0x6c43075e},
	"2601:4f8:221:42c2::2":                    {0x260104f8022142c2, 0x2},
	"2602:2350:5:104:bd:311c:7060:fd66":       {0x2602235000050104, 0xbd311c7060fd66},
	"2602:26f0:9100:865::2b0e":                {0x260226f091000865, 0x2b0e},
	"2607:6b90:7647:26b3::5:2107:6d01":        {0x26076b90764726b3, 0x521076d01},
	"2405:6e00:1fb6:5e00:906c:dbc6:5846:8bd":  {0x24056e001fb65e00, 0x906cdbc6584608bd},
	"2606:4700:7::629f:8955":                  {0x2606470000070000, 0x629f8955},
	"2600:1390::2:8556:3656:6382:d743":        {0x2600139000000002, 0x855636566382d743},
	"2600:d900:1::216:3eff:fe24:102":          {0x2600d90000010000, 0x2163efffe240102},
	"2602:2698:7c26:b18:f588:6253:4bcb:6644":  {0x260226987c260b18, 0xf58862534bcb6644},
	"2602:b80::6:35:6:0:1":                    {0x26020b8000000006, 0x35000600000001},
	"2003:cd:772f:1dc6:453e:c661:e9e:2e1e":    {0x200300cd772f1dc6, 0x453ec6610e9e2e1e},
	"2610:e0:6040:8b14:6d90:8e6:f087:64f9":    {0x261000e060408b14, 0x6d9008e6f08764f9},
	"2803:f800:50::6c62:c05c":                 {0x2803f80000500000, 0x6c62c05c},
	"2606:5686:45e8:249c:f66d:4ff:fe6d:ff52":  {0x2606568645e8249c, 0xf66d04fffe6dff52},
	"2601:16d:6600:160:9442:b7dc:6e61:4410":   {0x2601016d66000160, 0x9442b7dc6e614410},
	"2601:5ec0:2006:6898:e9ee:5462:e81e:c309": {0x26015ec020066898, 0xe9ee5462e81ec309},
	"2602:4780:8:687::2f20:1602:1":            {0x2602478000080687, 0x2f2016020001},
	"2400:3240:8909:1:1163:d3c2:2662:90e9":    {0x2400324089090001, 0x1163d3c2266290e9},
	"2405:204:418b:42dc::1bc8:b061":           {0x24050204418b42dc, 0x1bc8b061},
	"2600:9000:218f:1400:f:724f:6980:9361":    {0x26009000218f1400, 0xf724f69809361},
	"2606:4700:20::6816:d26":                  {0x2606470000200000, 0x68160d26},
	"2804:d41:434c:1600:d438:f46c:b88:4216":   {0x28040d41434c1600, 0xd438f46c0b884216},
	"2600:5fc1::120:45:153:59:199":            {0x26005fc100000120, 0x45015300590199},
	"2602:603f:696c:8700:1001:f6c6:916b:67fb": {0x2602603f696c8700, 0x1001f6c6916b67fb},
	"2602:6440:3286:1:7803:dc3f:4fdc:69d2":    {0x2602644032860001, 0x7803dc3f4fdc69d2},
	"2409:4064:516:9f0c:5516:261f:2d1c:6566":  {0x2409406405169f0c, 0x5516261f2d1c6566},
	"2800:110:2800:296:216:46ff:fe77:7122":    {0x2800011028000296, 0x21646fffe777122},
	"2600:1450:4007:80f::2013":                {0x260014504007080f, 0x2013},
	"2600:3c02::603c:9266:6e76:161d":          {0x26003c0200000000, 0x603c92666e76161d},
	"2001:8d8:100f:f000::216":                 {0x200108d8100ff000, 0x216},
	"2402:800:4188:cdd5:6536:5396:13f4:1d81":  {0x240208004188cdd5, 0x6536539613f41d81},
	"2601:249:8200:382:38dc:d633:34d9:b8f9":   {0x2601024982000382, 0x38dcd63334d9b8f9},
	"2602:4780:b:665::3201:7df4:1":            {0x26024780000b0665, 0x32017df40001},
	"2602:8109:6ec0:1034:d0f2:2669:fdd4:27f0": {0x260281096ec01034, 0xd0f22669fdd427f0},
	"2606:7d80:1:4::163:106:71":               {0x26067d8000010004, 0x16301060071},
	"2001:470:26:ef2::200":                    {0x2001047000260ef2, 0x200},
	"2600:9000:2117:6200:15:870:1e80:9361":    {0x2600900021176200, 0x1508701e809361},
	"2600:5fc1::120:45:153:59:203":            {0x26005fc100000120, 0x45015300590203},
	"2600:f940:2:1:2::702":                    {0x2600f94000020001, 0x2000000000702},
	"2602:2770:16::216:46ff:fe7d:5713":        {0x2602277000160000, 0x21646fffe7d5713},
	"2602:c207:2011:4247::1":                  {0x2602c20720114247, 0x1},
	"2606:5e80::":                             {0x26065e8000000000, 0x0},
	"2620:4d:4006:6259:7:2::1":                {0x2620004d40066259, 0x7000200000001},
	"2c06:6560:140::5c":                       {0x2c06656001400000, 0x5c},
	"2402:3680:1630:1056:7535:4c7c:e273:4d7c": {0x2402368016301056, 0x75354c7ce2734d7c},
	"2607:6400:2:15::19c":                     {0x2607640000020015, 0x19c},
	"2600:f940:4::47":                         {0x2600f94000040000, 0x47},
	"2602:c7d:3bbb:3800:e8d2:e07:26dc:9d6c":   {0x26020c7d3bbb3800, 0xe8d20e0726dc9d6c},
	"240e:62:c001:b1b9:d4c1:6350:4142:667":    {0x240e0062c001b1b9, 0xd4c1635041420667},
	"2409:4043:4e8c:6686:49c3:d737:b2eb:57b7": {0x240940434e8c6686, 0x49c3d737b2eb57b7},
	"2602:26f0:9100:16::6011:cec6":            {0x260226f091000016, 0x6011cec6},
	"2620:10f:5000:5002:96:124:243:2":         {0x2620010f50005002, 0x96012402430002},
	"240f:73:3c67:1:2856:e97c:dd41:db6e":      {0x240f00733c670001, 0x2856e97cdd41db6e},
	"2600:9000:2113:3400:1d:6eb4:1680:9361":   {0x2600900021133400, 0x1d6eb416809361},
	"2606:4700:10::6816:62d":                  {0x2606470000100000, 0x6816062d},
	"2607:6400:2:15::19f":                     {0x2607640000020015, 0x19f},
	"2001:fb1:136:d350:7461:29e2:43f7:6c54":   {0x20010fb10136d350, 0x746129e243f76c54},
	"2607:6b60:6:9000::20e1":                  {0x26076b6000069000, 0x20e1},
	"2001:16b8:246f:4800:5d6d:c0dd:5199:f249": {0x200116b8246f4800, 0x5d6dc0dd5199f249},
	"2001:6c8:25:789:c23:ff4:8db4:7":          {0x200106c800250789, 0xc230ff48db40007},
	"2400:cb00:2049:1::629f:16e6":             {0x2400cb0020490001, 0x629f16e6},
	"2402:800:61b3:84f7:f19e:f967:6163:7806":  {0x2402080061b384f7, 0xf19ef96761637806},
	"2409:8950:2c30:b0d2:1d23:e2b8:253e:86ce": {0x240989502c30b0d2, 0x1d23e2b8253e86ce},
	"2600:1158:1000:406::267":                 {0x2600115810000406, 0x267},
	"2600:367:2:60f::54":                      {0x260003670002060f, 0x54},
	"2402:4000:21c1:7b9d:65e3:b2c:744:634e":   {0x2402400021c17b9d, 0x65e30b2c0744634e},
	"2600:9000:2113:9000:17:2e62:cb40:9361":   {0x2600900021139000, 0x172e62cb409361},
	"2606:4700:20::6816:e63":                  {0x2606470000200000, 0x68160e63},
	"2601:540:42f2:5500:ecb2:d57:f9cc:e203":   {0x2601054042f25500, 0xecb20d57f9cce203},
	"2603:b0c0:3:f0::e6:0":                    {0x2603b0c0000300f0, 0xe60000},
	"2406:6c00::3407:133:16:199:19":           {0x24066c0000003407, 0x133001601990019},
	"2600:1f16:676f:4600:416:31fd:d6c9:5f13":  {0x26001f16676f4600, 0x41631fdd6c95f13},
	"2607:5501:3000:16b::2":                   {0x260755013000016b, 0x2},
	"2400:2410:9120:c400:746b:726c:6c71:196f": {0x240024109120c400, 0x746b726c6c71196f},
	"2606:4700::6611:bd4e":                    {0x2606470000000000, 0x6611bd4e},
	"2804:d45:361b:7600:f582:26c9:1718:2e65":  {0x28040d45361b7600, 0xf58226c917182e65},
	"2602:b80::6:35:6:0:2":                    {0x26020b8000000006, 0x35000600000002},
	"2600:9000:2113:6400:b:251:9680:9361":     {0x2600900021136400, 0xb025196809361},
	"2601:238:206:202:1068::":                 {0x2601023802060202, 0x1068000000000000},
	"2602:b18:581:1:5988:d457:b192:dc5b":      {0x26020b1805810001, 0x5988d457b192dc5b},
	"2001:6d6:105:1::1:0:5":                   {0x200106d601050001, 0x100000005},
	"2604:926:1::21":                          {0x2604092600010000, 0x21},
	"2606:2e00::1:225:9066:6e4e:4456":         {0x26062e0000000001, 0x22590666e4e4456},
	"2400:cb00:2049:1::629f:1b6d":             {0x2400cb0020490001, 0x629f1b6d},
	"2601:988:8200:3c66:6dbf:550f:26f2:e1e8":  {0x2601098882003c66, 0x6dbf550f26f2e1e8},
	"2601:4f9:3b:2692::2":                     {0x260104f9003b2692, 0x2},
	"2602:26f0:f4::17c6:e571":                 {0x260226f000f40000, 0x17c6e571},
	"2601:562:6603:3d40:7569:e360:1572:1104":  {0x2601056266033d40, 0x7569e36015721104},
	"2607:f1c0:fe:53:165:132:32:111":          {0x2607f1c000fe0053, 0x165013200320111},
	"2606:4700:20::6c43:4547":                 {0x2606470000200000, 0x6c434547},
	"2404:4404:231f:de00:5d63:d8c:db3d:903e":  {0x24044404231fde00, 0x5d630d8cdb3d903e},
	"2600:9000:219c:ec00:13:6662:d000:9361":   {0x26009000219cec00, 0x136662d0009361},
	"2602:26f0:b200::58dd:1892":               {0x260226f0b2000000, 0x58dd1892},
	"2001:f40:906:5f9e:f0ec:535e:1079:fb10":   {0x20010f4009065f9e, 0xf0ec535e1079fb10},
	"2600:1007:b123:2907:e630:f6b6:6626:9b79": {0x26001007b1232907, 0xe630f6b666269b79},
	"2001:d67:5300:2::10":                     {0x20010d6753000002, 0x10},
	"2400:80c0:1105:f4b:966:e662:df70:6963":   {0x240080c011050f4b, 0x966e662df706963},
	"2605:6440:3008:9000::64cc":               {0x2605644030089000, 0x64cc},
	"2606:4700:20::6c43:4412":                 {0x2606470000200000, 0x6c434412},
	"2600:5fc1::120:45:153:59:202":            {0x26005fc100000120, 0x45015300590202},
	"2606:82c2:0:1f6::1":                      {0x260682c2000001f6, 0x1},
	"2606:2b43:26c:d77c::":                    {0x26062b43026cd77c, 0x0},
	"2001:1662:6039:6400:1066:7611:7c11:46d7": {0x2001166260396400, 0x106676117c1146d7},
	"2001::4675:b261":                         {0x2001000000000000, 0x4675b261},
	"2607:fb90:90e6:c665:650b:f55d:20e:6d1":   {0x2607fb9090e6c665, 0x650bf55d020e06d1},
	"2600:d900:1::216:3eff:fe24:103":          {0x2600d90000010000, 0x2163efffe240103},
	"2600:3c03::603c:9266:6e60:4370":          {0x26003c0300000000, 0x603c92666e604370},
	"2003:eb:43d3:9ccd:7513:812f:b66c:bc91":   {0x200300eb43d39ccd, 0x7513812fb66cbc91},
	"2600:9000:218c:3600:1e:6cee:2fc0:9361":   {0x26009000218c3600, 0x1e6cee2fc09361},
	"2600:9000:219c:3400:1:97b0:cb40:9361":    {0x26009000219c3400, 0x197b0cb409361},
	"2602:908:1784:7480:45e1:18d9:80d6:b698":  {0x2602090817847480, 0x45e118d980d6b698},
	"2600:3c03::603c:9166:6eb0:dd3b":          {0x26003c0300000000, 0x603c91666eb0dd3b},
	"2001:981:b1cf:1:e1d0:6518:f76b:b9d9":     {0x20010981b1cf0001, 0xe1d06518f76bb9d9},
	"2600:1006:b06c:9f47:65ee:7b22:66fb:6496": {0x26001006b06c9f47, 0x65ee7b2266fb6496},
	"2400:6904::f03c:92ff:fee4:93b4":          {0x2400690400000000, 0xf03c92fffee493b4},
	"2606:4700:3030::6615:331":                {0x2606470030300000, 0x66150331},
	"2606:4700:e2::6c40:8404":                 {0x2606470000e20000, 0x6c408404},
	"2606:6000::":                             {0x2606600000000000, 0x0},
	"2804:14c:65d3:4196:2113:c662:e1ed:b8cc":  {0x2804014c65d34196, 0x2113c662e1edb8cc},
	"2600:23c4:2791:f800:d4d0:5cc6:6876:e308": {0x260023c42791f800, 0xd4d05cc66876e308},
	"2001:67c:2260:7060:5d56:5021:2e27:73de":  {0x2001067c22607060, 0x5d5650212e2773de},
	"2600:1f18:678f:4600:79b6:ec95:1bb3:5214": {0x26001f18678f4600, 0x79b6ec951bb35214},
	"2600:9000:2117:4400:1d:e6cc:9f40:9361":   {0x2600900021174400, 0x1de6cc9f409361},
	"2600:f10:401::1c00:c7ff:fe00:19c":        {0x26000f1004010000, 0x1c00c7fffe00019c},
	"2602:2350:5:101:80e2:1d6:fdec:e8c2":      {0x2602235000050101, 0x80e201d6fdece8c2},
	"2001:6d6:105:1::1:0:6":                   {0x200106d601050001, 0x100000006},
	"2601:40b:6402:6db3:5432:9d5e:942e:c257":  {0x2601040b64026db3, 0x54329d5e942ec257},
	"2601:53c0:ff04:ffff::65":                 {0x260153c0ff04ffff, 0x65},
	"2602:2698:4424:2e6c:5517:1848:b578:704":  {0x2602269844242e6c, 0x55171848b5780704},
	"2001:4453:369:b000:418c:b08e:fc45:b1b2":  {0x200144530369b000, 0x418cb08efc45b1b2},
	"2600:23c4:fb86:de00:3167:b016:196c:5ed6": {0x260023c4fb86de00, 0x3167b016196c5ed6},
	"2600:23c6:4c01:9e00:f479:13bc:e696:56ed": {0x260023c64c019e00, 0xf47913bce69656ed},
	"2600:7660:0:1097::1":                     {0x2600766000001097, 0x1},
	"2601:130:2000:118:89:146:248:231":        {0x2601013020000118, 0x89014602480231},
	"2601:4f8:242:1f16::3":                    {0x260104f802421f16, 0x3},
	"2602:c7f:1ecb:5d00:c557:6b7f:d501:9be2":  {0x26020c7f1ecb5d00, 0xc5576b7fd5019be2},
	"2602:f08:1000::":                         {0x26020f0810000000, 0x0},
	"2620:101:9060:53::55":                    {0x2620010190600053, 0x55},
	"2400:3200:2006:57::1":                    {0x2400320020060057, 0x1},
	"2001:1662:e6c1:9697:5c0e:716c:399d:16d0": {0x20011662e6c19697, 0x5c0e716c399d16d0},
	"2409:8628:3036:be50:15c0:71be:c523:cfc5": {0x240986283036be50, 0x15c071bec523cfc5},
	"2604:6000:8400:6100:6923:908b:2c96:d64f": {0x2604600084006100, 0x6923908b2c96d64f},
	"2600:b580:8000:12:d978:eff4:67ec:4537":   {0x2600b58080000012, 0xd978eff467ec4537},
	"2600:f940:5::190":                        {0x2600f94000050000, 0x190},
	"2601:e34:ec95:f450:308c:43d0:3159:b65":   {0x26010e34ec95f450, 0x308c43d031590b65},
	"2605:d014:9d6:8c10:624d:5e61:12e9:798b":  {0x2605d01409d68c10, 0x624d5e6112e9798b},
	"2606:98c1:50::6c40:2059":                 {0x260698c100500000, 0x6c402059},
	"2001:8003:3016:d501:6559:b658:1151:c1b1": {0x200180033016d501, 0x6559b6581151c1b1},
	"2003:d8:bf41:2900:bcd3:f0b6:d1d8:fb23":   {0x200300d8bf412900, 0xbcd3f0b6d1d8fb23},
	"2406:d614:57d:c801:7c2c:df71:f0d9:6fce":  {0x2406d614057dc801, 0x7c2cdf71f0d96fce},
	"240e:376:1b48:d900:476:5c47:3bbf:743e":   {0x240e03761b48d900, 0x4765c473bbf743e},
	"2600:9000:2111:1000:13:69dd:56c0:9361":   {0x2600900021111000, 0x1369dd56c09361},
	"2601:598:b009:d249:966:9d9e:496b:27b9":   {0x26010598b009d249, 0x9669d9e496b27b9},
	"2402:800:61b1:f076:c401:1bc0:8bef:3643":  {0x2402080061b1f076, 0xc4011bc08bef3643},
	"2607:fb90:429f:e897:932:10b3:fbff:6267":  {0x2607fb90429fe897, 0x93210b3fbff6267},
	"2602:c7f:d6c6:6200:9d65:c496:6763:4371":  {0x26020c7fd6c66200, 0x9d65c49667634371},
	"240e:66:9000:1100::196":                  {0x240e006690001100, 0x196},
	"2406:100::":                              {0x2406010000000000, 0x0},
}
