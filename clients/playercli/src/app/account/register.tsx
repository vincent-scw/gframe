import * as React from 'react';
import { SyntheticEvent } from 'react';
import { authService } from '../services';
import './register.scss';
import logo from '../../logo.png';

interface RegisterState {
	username: string;
	hasError: boolean;
	errorMsg: string;
}

export class Register extends React.Component<any, RegisterState> {
	constructor(props: any) {
		super(props);
		this.state = { username: '', hasError: false, errorMsg: '' };

		this.handleChange = this.handleChange.bind(this);
		this.submit = this.submit.bind(this);
	}

	handleChange(e: any) {
		this.setState({ username: e.target.value });
	}

	submit(e: SyntheticEvent) {
		e.preventDefault();
		if (this.state.username.trim() === '') {
			this.setState({ hasError: true, errorMsg: 'Username is required' })
			return;
		}
		authService.login(this.state.username);
	}

	render() {
		return (
			<div className="login">
				<img src={logo} className="logo" alt="logo" />
				<form onSubmit={this.submit}>
					<div className="field">
						<label className="label">Name</label>
						<div className="control has-icons-left">
							<input className="input" type="text" placeholder="Input Username"
								onChange={this.handleChange} />
							<span className="icon is-small is-left">
								<i className="fas fa-user"></i>
							</span>
						</div>
						{this.state.hasError &&
							<p className="help is-danger">{this.state.errorMsg}</p>
						}
					</div>
					<input className="button full-width" type="submit" value="Register" />
				</form>
			</div>
		);
	}
}