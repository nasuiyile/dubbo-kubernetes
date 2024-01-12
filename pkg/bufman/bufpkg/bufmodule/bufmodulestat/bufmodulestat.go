// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufmodulestat

import (
	"github.com/apache/dubbo-kubernetes/pkg/bufman/bufpkg/bufmodule"
	"github.com/apache/dubbo-kubernetes/pkg/bufman/pkg/protostat"
)

// NewFileWalker returns a new FileWalker for the given Module.
//
// This walks all target files from TargetFileInfos.
//
// We use TargetFileInfos instead of SourceFileInfos as this means
// that if someone sets up a filter at a higher level, this will respect it.
// In most cases, TargetFileInfos is the same as SourceFileInfos.
func NewFileWalker(module bufmodule.Module) protostat.FileWalker {
	return newFileWalker(module)
}
