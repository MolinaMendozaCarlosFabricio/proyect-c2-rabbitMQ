import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Requests } from '../models/request';
import { Observable } from 'rxjs';
import { Acquire } from '../models/acquire';
import { Message } from '../models/message';
import { MessageWhitResults } from '../models/message-whit-results';
import { RequestWhitStatus } from '../models/request-whit-status';

@Injectable({
  providedIn: 'root'
})
export class RequestsService {
  _url_api1:string = "http://localhost:8080/requests/"
  _url_api2:string = "http://localhost:9080/requests/"

  constructor(private _http: HttpClient) { }

  createRequest(request: Requests): Observable<MessageWhitResults<number>>{
    return this._http.post<MessageWhitResults<number>>(this._url_api1 + "request", request)
  }

  addProductToRequest(acquire: Acquire): Observable<Message>{
    return this._http.post<Message>(this._url_api1 + "product", acquire)
  }
  
  getRequestsMine(id_user: number): Observable<MessageWhitResults<RequestWhitStatus>>{
    return this._http.get<MessageWhitResults<RequestWhitStatus>>(this._url_api2 + id_user)
  }
}
