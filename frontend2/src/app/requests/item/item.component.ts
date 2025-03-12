import { Component, Input } from '@angular/core';
import { RequestWhitStatus } from '../../models/request-whit-status';

@Component({
  selector: 'app-item',
  templateUrl: './item.component.html',
  styleUrl: './item.component.css'
})
export class ItemComponent {
  @Input() request: RequestWhitStatus = {
    ID: 0,
    Date_request: new Date,
    Id_user: 0,
    Status: "Pendiente"
  };

}
