package main

var symbolTemplates = map[string]Symbols{
	"compatible": {
		Lock:          "RO",
		Network:       "SSH",
		Separator:     "\u25B6",
		SeparatorThin: "\u276F",

		RepoDetached:   "\u2693",
		RepoAhead:      "\u2B06",
		RepoBehind:     "\u2B07",
		RepoStaged:     "\u2714",
		RepoNotStaged:  "\u270E",
		RepoUntracked:  "+",
		RepoConflicted: "\u273C",
		RepoStashed:    "\u2691",
	},
	"patched": {
		Lock:          "\uE0A2",
		Network:       "\uE0A2",
		Separator:     "\uE0B0",
		SeparatorThin: "\uE0B1",

		RepoDetached:   "\u2693",
		RepoAhead:      "\u2B06",
		RepoBehind:     "\u2B07",
		RepoStaged:     "\u2714",
		RepoNotStaged:  "\u270E",
		RepoUntracked:  "+",
		RepoConflicted: "\u273C",
		RepoStashed:    "\u2691",
	},
	"flat": {
		RepoDetached:   "\u2693",
		RepoAhead:      "\u2B06",
		RepoBehind:     "\u2B07",
		RepoStaged:     "\u2714",
		RepoNotStaged:  "\u270E",
		RepoUntracked:  "+",
		RepoConflicted: "\u273C",
		RepoStashed:    "\u2691",
	},
}

var shellInfos = map[string]ShellInfo{
	"bash": {
		colorTemplate:    "\\[\\e%s\\]",
		rootIndicator:    "\\$",
		escapedBackslash: `\\\\`,
		escapedBacktick:  "\\`",
		escapedDollar:    `\$`,
	},
	"zsh": {
		colorTemplate:    "%%{\u001b%s%%}",
		rootIndicator:    "%#",
		escapedBackslash: `\\`,
		escapedBacktick:  "\\`",
		escapedDollar:    `\$`,
	},
	"bare": {
		colorTemplate: "%s",
		rootIndicator:    "$",
		escapedBackslash: `\`,
		escapedBacktick:  "`",
		escapedDollar:    `$`,
	},
}

var themes = map[string]Theme{
	"default": {
		Reset: 0xFF,

		UsernameFg:     250,
		UsernameBg:     240,
		UsernameRootBg: 124,

		HostnameFg: 250,
		HostnameBg: 238,

		HomeSpecialDisplay: true,
		HomeFg:             15,  // white
		HomeBg:             31,  // blueish
		PathFg:             250, // light grey
		PathBg:             237, // dark grey
		CwdFg:              254, // nearly-white grey
		SeparatorFg:        244,

		ReadonlyFg: 254,
		ReadonlyBg: 124,

		SshFg: 254,
		SshBg: 166, // medium orange

		DockerMachineFg: 177, // light purple
		DockerMachineBg: 55,  // purple

		KubeClusterFg:   117,
		KubeClusterBg:   26,
		KubeNamespaceFg: 170,
		KubeNamespaceBg: 17,

		DotEnvFg: 15, // white
		DotEnvBg: 55, // purple

		AWSFg: 15, // white
		AWSBg: 22, // dark green

		RepoCleanFg: 0,   // black
		RepoCleanBg: 148, // a light green color
		RepoDirtyFg: 15,  // white
		RepoDirtyBg: 161, // pink/red

		JobsFg: 39,
		JobsBg: 238,

		CmdPassedFg: 15,
		CmdPassedBg: 236,
		CmdFailedFg: 15,
		CmdFailedBg: 161,

		SvnChangesFg: 22, // dark green
		SvnChangesBg: 148,

		GitAheadFg:      250,
		GitAheadBg:      240,
		GitBehindFg:     250,
		GitBehindBg:     240,
		GitStagedFg:     15,
		GitStagedBg:     22,
		GitNotStagedFg:  15,
		GitNotStagedBg:  130,
		GitUntrackedFg:  15,
		GitUntrackedBg:  52,
		GitConflictedFg: 15,
		GitConflictedBg: 9,
		GitStashedFg:    15,
		GitStashedBg:    20,

		VirtualEnvFg: 00,
		VirtualEnvBg: 35, // a mid-tone green

		PerlbrewFg: 00,
		PerlbrewBg: 20, // a mid-tone blue

		TimeFg: 15,
		TimeBg: 236,

		ShellVarFg: 52,
		ShellVarBg: 11,

		HostnameColorizedFgMap: map[uint8]uint8{
			0:   250,
			1:   250,
			2:   120,
			3:   228,
			4:   250,
			5:   250,
			6:   123,
			7:   238,
			8:   0,
			9:   0,
			10:  0,
			11:  0,
			12:  250,
			13:  0,
			14:  0,
			15:  242,
			16:  250,
			17:  250,
			18:  250,
			19:  189,
			20:  254,
			21:  250,
			22:  83,
			23:  87,
			24:  117,
			25:  188,
			26:  254,
			27:  0,
			28:  120,
			29:  122,
			30:  123,
			31:  159,
			32:  255,
			33:  0,
			34:  157,
			35:  158,
			36:  159,
			37:  159,
			38:  195,
			39:  0,
			40:  194,
			41:  194,
			42:  195,
			43:  195,
			44:  195,
			45:  0,
			46:  0,
			47:  0,
			48:  0,
			49:  0,
			50:  0,
			51:  0,
			52:  250,
			53:  250,
			54:  250,
			55:  189,
			56:  254,
			57:  250,
			58:  227,
			59:  253,
			60:  255,
			61:  0,
			62:  233,
			63:  17,
			64:  192,
			65:  255,
			66:  195,
			67:  232,
			68:  233,
			69:  17,
			70:  193,
			71:  232,
			72:  232,
			73:  232,
			74:  234,
			75:  236,
			76:  194,
			77:  235,
			78:  235,
			79:  235,
			80:  235,
			81:  237,
			82:  0,
			83:  237,
			84:  237,
			85:  237,
			86:  237,
			87:  237,
			88:  250,
			89:  250,
			90:  250,
			91:  189,
			92:  254,
			93:  0,
			94:  222,
			95:  255,
			96:  255,
			97:  232,
			98:  233,
			99:  17,
			100: 228,
			101: 15,
			102: 232,
			103: 233,
			104: 17,
			105: 18,
			106: 229,
			107: 232,
			108: 234,
			109: 234,
			110: 236,
			111: 54,
			112: 230,
			113: 235,
			114: 22,
			115: 237,
			116: 238,
			117: 238,
			118: 0,
			119: 237,
			120: 22,
			121: 23,
			122: 23,
			123: 23,
			124: 252,
			125: 252,
			126: 189,
			127: 189,
			128: 254,
			129: 0,
			130: 223,
			131: 232,
			132: 232,
			133: 232,
			134: 233,
			135: 17,
			136: 229,
			137: 232,
			138: 233,
			139: 234,
			140: 53,
			141: 18,
			142: 229,
			143: 232,
			144: 234,
			145: 236,
			146: 17,
			147: 19,
			148: 230,
			149: 235,
			150: 238,
			151: 22,
			152: 23,
			153: 24,
			154: 0,
			155: 237,
			156: 22,
			157: 2,
			158: 29,
			159: 6,
			160: 254,
			161: 254,
			162: 254,
			163: 254,
			164: 254,
			165: 0,
			166: 255,
			167: 233,
			168: 233,
			169: 234,
			170: 234,
			171: 235,
			172: 230,
			173: 234,
			174: 52,
			175: 235,
			176: 53,
			177: 53,
			178: 230,
			179: 235,
			180: 236,
			181: 52,
			182: 53,
			183: 55,
			184: 230,
			185: 235,
			186: 238,
			187: 58,
			188: 240,
			189: 20,
			190: 0,
			191: 238,
			192: 58,
			193: 64,
			194: 35,
			195: 66,
			196: 0,
			197: 0,
			198: 0,
			199: 0,
			200: 0,
			201: 0,
			202: 0,
			203: 235,
			204: 235,
			205: 235,
			206: 235,
			207: 53,
			208: 0,
			209: 236,
			210: 52,
			211: 237,
			212: 53,
			213: 53,
			214: 0,
			215: 236,
			216: 238,
			217: 1,
			218: 89,
			219: 5,
			220: 0,
			221: 237,
			222: 58,
			223: 95,
			224: 131,
			225: 126,
			226: 0,
			227: 238,
			228: 58,
			229: 3,
			230: 143,
			231: 242,
			232: 250,
			233: 250,
			234: 250,
			235: 250,
			236: 250,
			237: 250,
			238: 251,
			239: 252,
			240: 188,
			241: 254,
			242: 254,
			243: 255,
			244: 0,
			245: 232,
			246: 233,
			247: 234,
			248: 235,
			249: 236,
			250: 237,
			251: 238,
			252: 239,
			253: 240,
			254: 242,
			255: 243,
		},
	},
	"low-contrast": {
		Reset: 0xFF,

		UsernameFg:     234,
		UsernameBg:     250,
		UsernameRootBg: 198,

		HostnameFg: 234,
		HostnameBg: 252,

		HomeSpecialDisplay: true,
		HomeFg:             30, // blueish-green
		HomeBg:             15, // white
		PathFg:             234,
		PathBg:             254,
		CwdFg:              236,
		SeparatorFg:        244,

		ReadonlyFg: 124,
		ReadonlyBg: 253,

		SshFg: 166, // medium orange
		SshBg: 254,

		DockerMachineFg: 55,  // purple
		DockerMachineBg: 177, // light purple

		KubeClusterFg:   117,
		KubeClusterBg:   26,
		KubeNamespaceFg: 170,
		KubeNamespaceBg: 17,

		DotEnvFg: 15, // white
		DotEnvBg: 55, // purple

		RepoCleanFg: 232, // black
		RepoCleanBg: 230, // light yellow
		RepoDirtyFg: 232, // black
		RepoDirtyBg: 223, // orange/peach

		JobsFg: 238,
		JobsBg: 39,

		CmdPassedFg: 18,
		CmdPassedBg: 7,
		CmdFailedFg: 254,
		CmdFailedBg: 124,

		SvnChangesFg: 148,
		SvnChangesBg: 22, // dark green

		GitAheadFg:      240,
		GitAheadBg:      250,
		GitBehindFg:     240,
		GitBehindBg:     251,
		GitStagedFg:     22,
		GitStagedBg:     15,
		GitNotStagedFg:  130,
		GitNotStagedBg:  15,
		GitUntrackedFg:  52,
		GitUntrackedBg:  15,
		GitConflictedFg: 9,
		GitConflictedBg: 15,
		GitStashedFg:    15,
		GitStashedBg:    20,

		VirtualEnvFg: 35, // a mid-tone green
		VirtualEnvBg: 254,

		PerlbrewFg: 20, // a mid-tone blue
		PerlbrewBg: 15,

		TimeFg: 236,
		TimeBg: 15,

		HostnameColorizedFgMap: map[uint8]uint8{
			0:   250,
			1:   250,
			2:   120,
			3:   228,
			4:   250,
			5:   250,
			6:   123,
			7:   238,
			8:   0,
			9:   0,
			10:  0,
			11:  0,
			12:  250,
			13:  0,
			14:  0,
			15:  242,
			16:  250,
			17:  250,
			18:  250,
			19:  189,
			20:  254,
			21:  250,
			22:  83,
			23:  87,
			24:  117,
			25:  188,
			26:  254,
			27:  0,
			28:  120,
			29:  122,
			30:  123,
			31:  159,
			32:  255,
			33:  0,
			34:  157,
			35:  158,
			36:  159,
			37:  159,
			38:  195,
			39:  0,
			40:  194,
			41:  194,
			42:  195,
			43:  195,
			44:  195,
			45:  0,
			46:  0,
			47:  0,
			48:  0,
			49:  0,
			50:  0,
			51:  0,
			52:  250,
			53:  250,
			54:  250,
			55:  189,
			56:  254,
			57:  250,
			58:  227,
			59:  253,
			60:  255,
			61:  0,
			62:  233,
			63:  17,
			64:  192,
			65:  255,
			66:  195,
			67:  232,
			68:  233,
			69:  17,
			70:  193,
			71:  232,
			72:  232,
			73:  232,
			74:  234,
			75:  236,
			76:  194,
			77:  235,
			78:  235,
			79:  235,
			80:  235,
			81:  237,
			82:  0,
			83:  237,
			84:  237,
			85:  237,
			86:  237,
			87:  237,
			88:  250,
			89:  250,
			90:  250,
			91:  189,
			92:  254,
			93:  0,
			94:  222,
			95:  255,
			96:  255,
			97:  232,
			98:  233,
			99:  17,
			100: 228,
			101: 15,
			102: 232,
			103: 233,
			104: 17,
			105: 18,
			106: 229,
			107: 232,
			108: 234,
			109: 234,
			110: 236,
			111: 54,
			112: 230,
			113: 235,
			114: 22,
			115: 237,
			116: 238,
			117: 238,
			118: 0,
			119: 237,
			120: 22,
			121: 23,
			122: 23,
			123: 23,
			124: 252,
			125: 252,
			126: 189,
			127: 189,
			128: 254,
			129: 0,
			130: 223,
			131: 232,
			132: 232,
			133: 232,
			134: 233,
			135: 17,
			136: 229,
			137: 232,
			138: 233,
			139: 234,
			140: 53,
			141: 18,
			142: 229,
			143: 232,
			144: 234,
			145: 236,
			146: 17,
			147: 19,
			148: 230,
			149: 235,
			150: 238,
			151: 22,
			152: 23,
			153: 24,
			154: 0,
			155: 237,
			156: 22,
			157: 2,
			158: 29,
			159: 6,
			160: 254,
			161: 254,
			162: 254,
			163: 254,
			164: 254,
			165: 0,
			166: 255,
			167: 233,
			168: 233,
			169: 234,
			170: 234,
			171: 235,
			172: 230,
			173: 234,
			174: 52,
			175: 235,
			176: 53,
			177: 53,
			178: 230,
			179: 235,
			180: 236,
			181: 52,
			182: 53,
			183: 55,
			184: 230,
			185: 235,
			186: 238,
			187: 58,
			188: 240,
			189: 20,
			190: 0,
			191: 238,
			192: 58,
			193: 64,
			194: 35,
			195: 66,
			196: 0,
			197: 0,
			198: 0,
			199: 0,
			200: 0,
			201: 0,
			202: 0,
			203: 235,
			204: 235,
			205: 235,
			206: 235,
			207: 53,
			208: 0,
			209: 236,
			210: 52,
			211: 237,
			212: 53,
			213: 53,
			214: 0,
			215: 236,
			216: 238,
			217: 1,
			218: 89,
			219: 5,
			220: 0,
			221: 237,
			222: 58,
			223: 95,
			224: 131,
			225: 126,
			226: 0,
			227: 238,
			228: 58,
			229: 3,
			230: 143,
			231: 242,
			232: 250,
			233: 250,
			234: 250,
			235: 250,
			236: 250,
			237: 250,
			238: 251,
			239: 252,
			240: 188,
			241: 254,
			242: 254,
			243: 255,
			244: 0,
			245: 232,
			246: 233,
			247: 234,
			248: 235,
			249: 236,
			250: 237,
			251: 238,
			252: 239,
			253: 240,
			254: 242,
			255: 243,
		},
	},
}
