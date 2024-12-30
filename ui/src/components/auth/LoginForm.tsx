import React, { useContext, useState } from 'react';
import {
	Box,
	Button,
	Divider,
	Link,
	TextField,
	Typography,
} from '@mui/material';
import { AuthContext } from '../../hooks/AuthContext';

const LoginForm: React.FC = () => {
	const [identifier, setIdentifier] = useState('');
	const [password, setPassword] = useState('');

	const { login } = useContext(AuthContext);

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		try {
			login(identifier, password);
		} catch (error) {
			console.error('Login failed:', error);
		}
	};

	return (
		<Box
			display="flex"
			justifyContent="center"
			alignItems="center"
			height="100vh">
			<Box width={400}>
				<Typography variant="h4" align="center" gutterBottom>
					Sign in
				</Typography>
				<Box component="form" onSubmit={handleSubmit}>
					<TextField
						label="Username or Email"
						variant="outlined"
						fullWidth
						margin="normal"
						onChange={(e) => setIdentifier(e.target.value)}
					/>
					<TextField
						label="Password"
						type="password"
						variant="outlined"
						fullWidth
						margin="normal"
						onChange={(e) => setPassword(e.target.value)}
					/>
					<Button
						type="submit"
						variant="contained"
						color="primary"
						fullWidth
						sx={{ mt: 2, py: 1.5 }}>
						Sign in
					</Button>
				</Box>

				<Divider sx={{ my: 3 }}>OR</Divider>

				<Typography variant="body2" align="center">
					Don't have an account?{' '}
					<Link href="/signup" color="primary" underline="hover">
						Sign up
					</Link>
				</Typography>
			</Box>
		</Box>
	);
};

export default LoginForm;
