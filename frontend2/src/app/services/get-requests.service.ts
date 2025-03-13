import { Injectable } from '@angular/core';
import { io, Socket } from 'socket.io-client';
import { MessageWhitResults } from '../models/message-whit-results';
import { RequestWhitStatus } from '../models/request-whit-status';
import { Observable } from 'rxjs';
import { RabbitmqMessage } from '../models/rabbitmq-message';

@Injectable({
  providedIn: 'root'
})
export class GetRequestsService {
  _url: string = "http://localhost:7080"
  private socket: Socket;

  constructor() {
    this.socket = io(this._url, {
      path: "/socket.io/",
      withCredentials: true,
    });
  }

  onMessage(): Observable<RabbitmqMessage> {
    return new Observable((observer) => {
      this.socket.on('rabbitmq_message', (msg: RabbitmqMessage) => {
        observer.next(msg);
      });

      return () => this.socket.disconnect();
    });
  }
}
