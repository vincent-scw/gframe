import * as React from 'react';
import { SyntheticEvent } from 'react';
import axios from "axios";

interface User {
	username: string;
}

export class Register extends React.Component<any, User> {
	constructor(props: any) {
		super(props);
		this.state = { username: '' };

		this.handleChange = this.handleChange.bind(this);
		this.submit = this.submit.bind(this);
	}

	handleChange(e: any) {
		this.setState({ username: e.target.value });
	}

	submit(e: SyntheticEvent) {
		let data = new FormData();
		data.set("client_id", "player_api");
		data.set("client_secret", "999999");
		data.set("grant_type", "password");
		data.set("username", this.state.username);
		data.set("password", "123");
		axios.post("http://localhost:9096/token", data)
			.then(res => { console.log(res) })
			.catch(err => console.error(err));
		e.preventDefault();
	}

	render() {
		return (
			<section className="section">
				<form onSubmit={this.submit}>
					<h1 className="title">Register</h1>
					<div className="field">
						<div className="control">
							<input className="input is-primary" type="text" placeholder="Input Username"
								onChange={this.handleChange} />
						</div>
					</div>
					<input className="button is-info" type="submit" value="Submit" />
				</form>
			</section>
		);
	}
}