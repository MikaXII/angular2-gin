import { Component, OnInit, NgModule } from '@angular/core';
import { DataService } from './service/sample.service';
import { MsgSample } from "./models/sample.model"
import { Configuration} from "./service/app.constants"
import { BrowserModule } from '@angular/platform-browser';


@NgModule({
    imports: [BrowserModule]
})

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [DataService, Configuration]
})

export class AppComponent implements OnInit {
  //title = 'app works!';
  public title: string;
  constructor(private _dataService: DataService){}

  ngOnInit(){
    this.getPing();
    
  }

  private getPing(): void {
      this._dataService
      .GetPing()
      .subscribe((data:MsgSample) => this.title = data.message,
      error => console.log(error),
      () => console.log("pong") );
  }
}
