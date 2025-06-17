// Copyright © 2025 Luka Ivanović
// This code is licensed under the terms of the MIT licence (see LICENCE for details)

package main

type Program struct {
	Name     string
	External bool
	Command  string
}

var programs []Program

func init() {
	programs = []Program{
		{"Toggle theme", true, ",toggle-theme"},
		{"Start torrenting daemon", true, ",torrenting_start"},
		{"Stop torrenting daemon", true, ",torrenting_stop"},
		{"Fix adb", true, ",adb-fix"},
		{"Start tweety", true, "twtty up"},
		{"Stop tweety", true, "twtty down"},
		{"Share with QR", true, ",share"},
		{"Get status info", true, ",da -notick"},
		{"Enter emoji", true, "rofimoji"},
	}
}
