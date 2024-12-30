import {
	Box,
	TextField,
	Button,
	Typography,
	Link,
	Divider,
} from '@mui/material';

const Register = () => {
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
				<Typography
					variant="h4"
					fontWeight="bold"
					align="center"
					gutterBottom>
					Sign up
				</Typography>
				<Box component="form" noValidate autoComplete="off">
					<TextField
						label="Name"
						type="text"
						variant="outlined"
						fullWidth
						margin="normal"
					/>
					<TextField
						label="Email"
						type="email"
						variant="outlined"
						fullWidth
						margin="normal"
					/>
					<TextField
						label="Password"
						type="password"
						variant="outlined"
						fullWidth
						margin="normal"
					/>
					<Button
						type="submit"
						variant="contained"
						color="primary"
						fullWidth
						sx={{ mt: 2, py: 1.5 }}>
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
			</Box>
		</Box>
	);
};

export default Register;
