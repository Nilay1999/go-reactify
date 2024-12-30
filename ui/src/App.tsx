import { useContext } from 'react';
import {
	BrowserRouter as Router,
	Route,
	Routes,
	Navigate,
	Outlet,
} from 'react-router-dom';
import { AuthProvider, AuthContext } from './hooks/AuthContext';
import Register from './components/auth/Register';
import LoginPage from './components/auth/LoginForm';
import { Dashboard } from './components/Dashboard';

const App = () => {
	const PrivateRoute = () => {
		const { authenticated } = useContext(AuthContext);
		if (!authenticated) return <Navigate to="/login" replace />;

		return <Outlet />;
	};

	return (
		<AuthProvider>
			<Router>
				<Routes>
					<Route path="/login" element={<LoginPage />} />
					<Route path="/signup" element={<Register />} />
					<Route element={<PrivateRoute />}>
						<Route path="/" element={<Dashboard />} />
					</Route>
				</Routes>
			</Router>
		</AuthProvider>
	);
};

export default App;
