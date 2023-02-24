import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { ComponentsModule } from 'src/common/components/components.module';
import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import { EncyclopediaPageModule } from './pages/encyclopedia/encyclopedia.module';
import { MarketComponent } from './pages/market/market.component';

@NgModule({
	declarations: [AppComponent, MarketComponent],
	imports: [
		BrowserModule,
		HttpClientModule,
		RouterModule,
		AppRoutingModule,
		ComponentsModule,
		ReactiveFormsModule,
		EncyclopediaPageModule,
	],
	providers: [],
	bootstrap: [AppComponent],
})
export class AppModule {}
