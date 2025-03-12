import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MainPageComponent } from './main-page/main-page.component';
import { AddFormComponent } from './add-form/add-form.component';
import { ItemComponent } from './item/item.component';
import { HttpClientModule } from '@angular/common/http';



@NgModule({
  declarations: [
    MainPageComponent,
    AddFormComponent,
    ItemComponent
  ],
  imports: [
    CommonModule,
    HttpClientModule
  ],
  exports: [
    MainPageComponent
  ]
})
export class RequestsModule { }
