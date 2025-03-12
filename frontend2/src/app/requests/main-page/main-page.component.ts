import { Component, OnInit } from '@angular/core';
import { RequestWhitStatus } from '../../models/request-whit-status';
import { GetRequestsService } from '../../services/get-requests.service';

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrl: './main-page.component.css'
})
export class MainPageComponent implements OnInit{
  request_list: RequestWhitStatus[] = [];
  isLoading: boolean = true;

  constructor(private getRequestServices: GetRequestsService){}

  ngOnInit(): void {
    /*
    console.log("Conectando con websocket")
    this.getRequestServices.get_requests(1, (requests) => {
      this.request_list = requests.Results
      console.log("Resultados:", requests.Results)
    })
    console.log("Resultados:", this.request_list)
    this.getRequestServices.on_new_request((request) => this.request_list.push(request.Results[0]))
    */
  }
}
