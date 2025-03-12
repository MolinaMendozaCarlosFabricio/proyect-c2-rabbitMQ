import { Injectable } from '@angular/core';
import { io, Socket } from 'socket.io-client';
import { MessageWhitResults } from '../models/message-whit-results';
import { RequestWhitStatus } from '../models/request-whit-status';

@Injectable({
  providedIn: 'root'
})
export class GetRequestsService {
  _url: string = "http://localhost:7080"
  private socket: Socket;

  constructor() {
    this.socket = io(this._url, {
      transports: ['websocket'], // Fuerza el uso de WebSockets
      withCredentials: true
    });

    this.socket.on('connect', () => {
      console.log('Conectado a WebSocket');
    });
  }

  get_requests(id_user: number, callback: (response: MessageWhitResults<RequestWhitStatus>) => void){
    this.socket.emit('get_requests', id_user);
    this.socket.on('requests_list', callback)
  }

  send_request(id_request: number){
    this.socket.emit("add_request", id_request)
  }

  on_new_request(callback: (response: MessageWhitResults<RequestWhitStatus>) => void){
    this.socket.on("one_request", callback)
  }
}
