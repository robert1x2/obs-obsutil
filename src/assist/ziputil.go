// Copyright 2019 Huawei Technologies Co.,Ltd.
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use
// this file except in compliance with the License.  You may obtain a copy of the
// License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations under the License.
package assist

import (
	"archive/zip"
	"io"
	"os"
)

func DoCompress(file *os.File, prefix string, zw *zip.Writer) (int64, error) {
	info, err := file.Stat()
	if err != nil {
		return 0, err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return 0, err
	}
	header.Method = zip.Deflate
	if prefix != "" {
		header.Name = prefix + "/" + header.Name
	}
	fw, err := zw.CreateHeader(header)
	if err != nil {
		return 0, err
	}
	wcnt, err := io.Copy(fw, file)
	if err != nil {
		return 0, err
	}
	return wcnt, nil
}
