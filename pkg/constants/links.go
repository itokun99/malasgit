package constants

type Docs struct {
	CustomPagers      string
	CustomCommands    string
	CustomKeybindings string
	Keybindings       string
	Undoing           string
	Config            string
	Tutorial          string
	CustomPatchDemo   string
}

var Links = struct {
	Docs        Docs
	Issues      string
	Donate      string
	Discussions string
	RepoUrl     string
	Releases    string
}{
	RepoUrl:     "https://github.com/itokun99/malasgit",
	Issues:      "https://github.com/itokun99/malasgit/issues",
	Donate:      "https://github.com/sponsors/jesseduffield",
	Discussions: "https://github.com/itokun99/malasgit/discussions",
	Releases:    "https://github.com/itokun99/malasgit/releases",
	Docs: Docs{
		CustomPagers:      "https://github.com/itokun99/malasgit/blob/master/docs/Custom_Pagers.md",
		CustomKeybindings: "https://github.com/itokun99/malasgit/blob/master/docs/keybindings/Custom_Keybindings.md",
		CustomCommands:    "https://github.com/itokun99/malasgit/wiki/Custom-Commands-Compendium",
		Keybindings:       "https://github.com/itokun99/malasgit/blob/%s/docs/keybindings",
		Undoing:           "https://github.com/itokun99/malasgit/blob/master/docs/Undoing.md",
		Config:            "https://github.com/itokun99/malasgit/blob/%s/docs/Config.md",
		Tutorial:          "https://youtu.be/VDXvbHZYeKY",
		CustomPatchDemo:   "https://github.com/itokun99/malasgit#rebase-magic-custom-patches",
	},
}
