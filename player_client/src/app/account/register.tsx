import * as React from 'react';
import { SyntheticEvent } from 'react';
import User from './user.model';
import userService from '../services/auth.service';

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
		userService.login(this.state.username);
		e.preventDefault();
	}

	render() {
		return (
			<section className="section">
				<form onSubmit={this.submit}>
					<div className="field">
						<div className="control">
							<input className="input is-primary" type="text" placeholder="Input Username"
								onChange={this.handleChange} />
						</div>
					</div>
					<input className="button is-info" type="submit" value="Register" />
				</form>
			</section>
		);
	}
}