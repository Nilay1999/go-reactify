import {
	Box,
	TextField,
	Button,
	Typography,
	Link,
	Divider,
	FormControl,
	FormControlLabel,
	Radio,
	RadioGroup,
	FormLabel,
} from '@mui/material';
import { useState } from 'react';
import { signup } from './types/auth.types';

const Register = () => {
	const [username, setUsername] = useState<string>('');
	const [email, setEmail] = useState<string>('');
	const [password, setPassword] = useState<string>('');
	const [age, setAge] = useState<number>(0);
	const [gender, setGender] = useState<string>('');
	const [loading, setLoading] = useState<boolean>(false);
	const [error, setError] = useState<string | null>(null);

	const handleChangeAge = (e: React.ChangeEvent<HTMLInputElement>) => {
		setAge(Number(e.target.value));
	};

	const handleChangeUsername = (e: React.ChangeEvent<HTMLInputElement>) => {
		setUsername(e.target.value);
	};

	const handleChangeEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
		setEmail(e.target.value);
	};

	const handleChangeGender = (e: React.ChangeEvent<HTMLInputElement>) => {
		setGender(e.target.value);
	};

	const handleChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
		setPassword(e.target.value);
	};

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();

		const payload: signup = {
			age,
			email,
			gender,
			password,
			username,
		};

		setLoading(true);
		setError(null);

		try {
		} catch (err) {
			setError(null);
		} finally {
			setLoading(false);
		}
	};

	return (
		<Box
			sx={{
				display: 'flex',
				justifyContent: 'center',
				alignItems: 'center',
				height: '100vh',
				backgroundColor: '#f5f5f5',
			}}>
			<Box
				sx={{
					width: 400,
					p: 4,
					borderRadius: 2,
					boxShadow: '0px 4px 10px rgba(0, 0, 0, 0.1)',
					backgroundColor: '#fff',
				}}>
				<Typography variant="h4" align="center" gutterBottom>
					Welcome
				</Typography>
				<Typography
					variant="body2"
					align="center"
					color="textSecondary"
					mb={2}>
					Sign up to continue
				</Typography>
				<FormControl>
					<Box component="form" noValidate autoComplete="off">
						<TextField
							label="Username"
							type="text"
							variant="outlined"
							fullWidth
							margin="normal"
							required
							onChange={handleChangeUsername}
						/>
						<TextField
							label="Email"
							type="email"
							variant="outlined"
							fullWidth
							margin="normal"
							required
							onChange={handleChangeEmail}
						/>
						<TextField
							label="Password"
							type="password"
							variant="outlined"
							fullWidth
							margin="normal"
							required
							onChange={handleChangePassword}
						/>

						<TextField
							label="Age"
							type="number"
							variant="outlined"
							fullWidth
							margin="normal"
							required
							onChange={handleChangeAge}
						/>
						<RadioGroup
							row
							sx={{
								display: 'flex',
								justifyContent: 'center',
								alignItems: 'center',
							}}>
							<FormLabel sx={{ mr: 3 }}>Gender</FormLabel>
							<FormControlLabel
								value="female"
								control={
									<Radio onChange={handleChangeGender} />
								}
								label="Female"
							/>
							<FormControlLabel
								value="male"
								control={
									<Radio onChange={handleChangeGender} />
								}
								label="Male"
							/>
						</RadioGroup>
						<Button
							type="submit"
							variant="contained"
							color="primary"
							fullWidth
							sx={{ mt: 2, py: 1.5 }}
							onSubmit={handleSubmit}>
							Sign up
						</Button>
					</Box>

					<Divider sx={{ my: 3 }}>OR</Divider>

					<Typography variant="body2" align="center">
						Already have an account?{' '}
						<Link href="/login" color="primary" underline="hover">
							Sign in
						</Link>
					</Typography>
				</FormControl>
			</Box>
		</Box>
	);
};

export default Register;
