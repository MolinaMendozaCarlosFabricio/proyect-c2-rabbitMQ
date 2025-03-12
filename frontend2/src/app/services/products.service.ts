import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { MessageWhitResults } from '../models/message-whit-results';
import { Product } from '../models/product';
import { Message } from '../models/message';

@Injectable({
  providedIn: 'root'
})
export class ProductsService {
  _url_api2: string = "http://localhost:9080/products/"

  constructor(private _http: HttpClient) { }

  createProduct(product: Product): Observable<Message>{
    return this._http.post<Message>(this._url_api2, product)
  }

  getProducts(): Observable<MessageWhitResults<Product>>{
    return this._http.get<MessageWhitResults<Product>>(this._url_api2)
  }

  editProduct(product: Product): Observable<Message>{
    return this._http.put<Message>(this._url_api2 + product.ID, product)
  }

  deleteProduct(id_product: number): Observable<Message>{
    return this._http.delete<Message>(this._url_api2 + id_product)
  }

  getProductsOfRequest(id_request: number): Observable<MessageWhitResults<Product>>{
    return this._http.get<MessageWhitResults<Product>>(this._url_api2 + "request/" + id_request)
  }
}
