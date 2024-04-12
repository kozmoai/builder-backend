// Copyright 2024 Kozmoai Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filter

import (
	"errors"

	"github.com/kozmoai/builder-backend/src/websocket"
)

func SignalLeave(hub *websocket.Hub, message *websocket.Message) error {
	currentClient, hit := hub.Clients[message.ClientID]
	if !hit {
		return errors.New("[SignalLeave] target client(" + message.ClientID.String() + ") does dot exists.")
	}

	// broadcast in room users
	inRoomUsers := hub.GetInRoomUsersByRoomID(currentClient.APPID)
	inRoomUsers.LeaveRoom(currentClient.ExportMappedUserIDToString())
	message.SetBroadcastType(websocket.BROADCAST_TYPE_ENTER)
	message.RewriteBroadcast()
	message.SetBroadcastPayload(inRoomUsers.FetchAllInRoomUsers())
	hub.BroadcastToOtherClients(message, currentClient)

	// broadcast attached components users
	message.SetBroadcastType(websocket.BROADCAST_TYPE_ATTACH_COMPONENT)
	message.RewriteBroadcast()
	message.SetBroadcastPayload(inRoomUsers.FetchAllAttachedUsers())
	hub.BroadcastToOtherClients(message, currentClient)

	// kick leaved user
	hub.KickClient(currentClient)
	hub.CleanRoom(currentClient.APPID)
	return nil
}
