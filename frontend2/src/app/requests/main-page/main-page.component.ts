import { Component, OnInit } from '@angular/core';
import { RequestWhitStatus } from '../../models/request-whit-status';
import { GetRequestsService } from '../../services/get-requests.service';
import { RequestsService } from '../../services/requests.service';
import { RabbitmqMessage } from '../../models/rabbitmq-message';

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrl: './main-page.component.css'
})
export class MainPageComponent implements OnInit{
  messages: RabbitmqMessage[] = [];
  list_requests: RequestWhitStatus[] = [];

  constructor(
    private getRequestService: GetRequestsService,
    private requestService: RequestsService
  ) {}

  ngOnInit(): void {
    setInterval(() => this.polling(), 5000)
    /*
    this.getRequestService.onMessage().subscribe((msg) => {
      this.messages.push(msg);
    });
*/
  }

  polling(){
    this.requestService.getRequestsMine(1).subscribe(
      response => {
        console.log("Respuesta recibida:", response)
        this.list_requests = response.Results;
      },
      error => console.log("Error:", error)
    )
  }
}
