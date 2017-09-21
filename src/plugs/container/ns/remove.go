// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ns

import (
	log "github.com/cihub/seelog"
	docker "github.com/fsouza/go-dockerclient"
)

func Remove(client *docker.Client, container *docker.Container) error {
	log.Debugf("Removing container %s", container.ID)
	err := client.RemoveContainer(docker.RemoveContainerOptions{
		ID:    container.ID,
		Force: true,
	})
	log.Debugf("Removed container %s", container.ID)
	if err != nil {
		log.Error(err)
	}
	return err
}
