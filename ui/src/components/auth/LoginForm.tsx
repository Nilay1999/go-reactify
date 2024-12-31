import React, { useContext, useState } from 'react';
import {
	Alert,
	Box,
	Button,
	CircularProgress,
	Divider,
	Link,
	TextField,
	Typography,
} from '@mui/material';
import { AuthContext } from '../../hooks/AuthContext';

const LoginForm: React.FC = () => {
	const [identifier, setIdentifier] = useState<string>('');
	const [password, setPassword] = useState<string>('');
	const [loading, setLoading] = useState<boolean>(false);
	const [error, setError] = useState<string | null>('');

	const { login } = useContext(AuthContext);

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		setLoading(true);
		setError(null);

		try {
			await login(identifier, password);
		} catch (err) {
			setError('Invalid credentials or an error occurred.');
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
				width={400}
				p={4}
				borderRadius={2}
				boxShadow={3}
				bgcolor="white">
				<Typography variant="h4" align="center" gutterBottom>
					Welcome Back
				</Typography>
				<Typography
					variant="body2"
					align="center"
					color="textSecondary"
					mb={2}>
					Sign in to continue
				</Typography>

				<Box
					component="form"
					onSubmit={handleSubmit}
					sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
					<TextField
						label="Username or Email"
						variant="outlined"
						fullWidth
						onChange={(e) => setIdentifier(e.target.value)}
						onFocus={() => setError('')}
						autoFocus
						required
					/>
					<TextField
						label="Password"
						type="password"
						variant="outlined"
						fullWidth
						onChange={(e) => setPassword(e.target.value)}
						onFocus={() => setError('')}
						required
					/>
					<Button
						type="submit"
						variant="contained"
						color="primary"
						fullWidth
						disabled={loading}
						sx={{ py: 1.5 }}>
						{loading ? (
							<CircularProgress
								size={24}
								sx={{ color: 'white' }}
							/>
						) : (
							'Sign In'
						)}
					</Button>
				</Box>

				<Divider sx={{ my: 3 }}>OR</Divider>

				<Typography
					variant="body2"
					align="center"
					color="textSecondary">
					Don't have an account?{' '}
					<Link href="/signup" color="primary" underline="hover">
						Sign up
					</Link>
				</Typography>
				{error && (
					<Alert sx={{ mt: 2 }} severity="info">
						{error}
					</Alert>
				)}
			</Box>
		</Box>
	);
};

export default LoginForm;
