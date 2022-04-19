/***** BEGIN LICENSE BLOCK *****
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this file,
# You can obtain one at http://mozilla.org/MPL/2.0/.
#
# The Initial Developer of the Original Code is the Mozilla Foundation.
# Portions created by the Initial Developer are Copyright (C) 2012
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Rob Miller (rmiller@mozilla.com)
#
# ***** END LICENSE BLOCK *****/

package test

import (
	"testing"

	"github.com/rafrombrc/gospec/src/gospec"
	// payload_encoder "github.com/wulog/plugin-payload_encoder"
	// rst_encoder "github.com/wulog/plugin-rst_encoder"
	// scribble_decoder "github.com/wulog/plugin-scribble_decoder"
)

func TestAllSpecs(t *testing.T) {
	r := gospec.NewRunner()
	r.Parallel = false

	r.AddSpec(LoadFromConfigSpec)
	// r.AddSpec(scribble_decoder.ScribbleDecoderSpec) // todo go 不支持导出测试方法，优化 test
	// r.AddSpec(payload_encoder.PayloadEncoderSpec)
	// r.AddSpec(rst_encoder.RstEncoderSpec)

	gospec.MainGoTest(r, t)
}
