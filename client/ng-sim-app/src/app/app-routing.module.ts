import { NgModule } from '@angular/core';
import { Route, RouterModule, Routes } from '@angular/router';

const encyclopediaRouteConfig: Route = {
	loadChildren: () =>
		import('./components/encyclopedia/encyclopedia.module').then(
			(m) => m.EncyclopediaModule
		),
};

const routes: Routes = [
	{
		path: '',
		children: [
			{ path: '', ...encyclopediaRouteConfig },
			{ path: 'encyclopedia', ...encyclopediaRouteConfig },
		],
	},
];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule],
})
export class AppRoutingModule {}
