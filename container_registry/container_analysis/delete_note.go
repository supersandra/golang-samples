// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// [START containeranalysis_delete_note]

import (
	"context"
	"fmt"

	containeranalysis "cloud.google.com/go/containeranalysis/apiv1"
	grafeaspb "google.golang.org/genproto/googleapis/grafeas/v1"
)

// deleteNote removes an existing Note from the server.
func deleteNote(noteID, projectID string) error {
	// noteID := fmt.Sprintf("my-note")
	ctx := context.Background()
	client, err := containeranalysis.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("NewGrafeasV1Beta1Client: %v", err)
	}
	defer client.Close()

	req := &grafeaspb.DeleteNoteRequest{
		Name: fmt.Sprintf("projects/%s/notes/%s", projectID, noteID),
	}
	return client.GetGrafeasClient().DeleteNote(ctx, req)
}

// [END containeranalysis_delete_note]
