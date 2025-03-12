import { Component, OnInit } from '@angular/core';
import { Requests } from '../../models/request';
import { Product } from '../../models/product';
import { Acquire } from '../../models/acquire';
import { RequestsService } from '../../services/requests.service';
import { ProductsService } from '../../services/products.service';
import { GetRequestsService } from '../../services/get-requests.service';

@Component({
  selector: 'app-add-form',
  templateUrl: './add-form.component.html',
  styleUrl: './add-form.component.css'
})
export class AddFormComponent implements OnInit{
  new_request: Requests = {
    ID: 0,
    Id_user: 1,
    Id_status: 3,
    Date_request: new Date
  }
  new_id: number = 0;

  product_list: Product[] = [];

  acquires: Acquire[] = [];

  constructor(
    private requestServices: RequestsService, 
    private productServices: ProductsService,
    private getRequestServices: GetRequestsService
  ){}

  ngOnInit(): void {
    this.productServices.getProducts().subscribe(
      response => {
        console.log("Respuesta correcta del servidor")
        this.product_list = response.Results
      },
      error => console.log("Error:", error)
    )
}

create(){
  console.log("Creando pedido")
  this.requestServices.createRequest(this.new_request).subscribe(
    (response) => {
      this.new_id = response.Results[0];
      console.log("Id del pedido:", this.new_id);
  
      // Ahora añade los productos
      this.acquires.forEach(acquire => {
        acquire.Id_request = this.new_id;
        this.requestServices.addProductToRequest(acquire).subscribe(
          (res) => console.log("Producto añadido:", res),
          (err) => console.error("Error al añadir producto:", err)
        );
      });
  
      //this.getRequestServices.send_request(this.new_id);
    },
    (error) => console.error("Error al crear pedido:", error)
  );

  this.new_request = {
    ID: 0,
    Id_user: 1,
    Id_status: 3,
    Date_request: new Date
  }

  this.new_id = 0
  }

  addProduct(id_product: number){
    let new_acquire: Acquire = {
      Id_product: id_product,
      Id_request: 0,
      Quantity: 1
    };

    console.log("Agregando producto")

    this.acquires.push(new_acquire)
}
}
