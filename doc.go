// Copyright 2020 zephinzer. Use of this source code
// is governed by the MIT license that can be found
// in the LICENSE file

/*
github.com/usvc/go-config is a package to manage configuration
values in libraries/applications.

Defining a configuration map:

	var conf = config.Map{
		"bool_key": &config.Bool{},
		"float_key": &config.Float{},
		"int_key": &config.Int{},
		"int_slice_key": &config.IntSlice{},
		"string_key": &config.String{},
		"uint_key": &config.Uint{},
		"uint_slice_key": &config.UintSlice{},
	}

Loading environment variables (assuming conf is defined):

	conf.LoadFromEnvironment()

Applying to an instance of cobra.Command:

	command := cobra.Command{ ... }
	conf.ApplyToCobra(&command)

Applying to an instance of cobra.Command persistently:

	command := cobra.Command{ ... }
	conf.ApplyToCobraPersistent(&command)
*/
package config
