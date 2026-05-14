import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { AuthProvider } from './contexts/AuthContext';
import { ThemeProvider } from './contexts/ThemeContext';
import { ToastProvider } from './contexts/ToastContext';
import ToastContainer from './components/UI/ToastContainer';
import Header from './components/Layout/Header';
import BurgerMenu from './components/Layout/BurgerMenu';
import Sidebar from './components/Layout/Sidebar';
import ProtectedRoute from './components/Layout/ProtectedRoute';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import Cards from './pages/Cards';
import Terminals from './pages/Terminals';
import Transactions from './pages/Transactions';
import Keys from './pages/Keys';
import Profile from './pages/Profile';
import TerminalTest from './pages/TerminalTest';
import './index.css';

function Layout({ children, showBurger = true }) {
  return (
    <div className="app-layout">
      <Header showBurger={showBurger} />
      <div className="app-content">
        <Sidebar />
        <main className="app-main">{children}</main>
      </div>
    </div>
  );
}

function AppRoutes() {
  return (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      
      <Route path="/" element={
        <ProtectedRoute>
          <Layout><Dashboard /></Layout>
        </ProtectedRoute>
      } />
      
      <Route path="/cards" element={
        <ProtectedRoute>
          <Layout><Cards /></Layout>
        </ProtectedRoute>
      } />
      
      <Route path="/terminals" element={
        <ProtectedRoute>
          <Layout><Terminals /></Layout>
        </ProtectedRoute>
      } />
      
      <Route path="/transactions" element={
        <ProtectedRoute>
          <Layout><Transactions /></Layout>
        </ProtectedRoute>
      } />
      
      <Route path="/keys" element={
        <ProtectedRoute adminOnly>
          <Layout><Keys /></Layout>
        </ProtectedRoute>
      } />
      
      <Route path="/profile" element={
        <ProtectedRoute>
          <Layout><Profile /></Layout>
        </ProtectedRoute>
      } />
      
      <Route path="/terminal-test" element={
        <ProtectedRoute>
          <Layout><TerminalTest /></Layout>
        </ProtectedRoute>
      } />
      
      <Route path="*" element={<Navigate to="/" />} />
    </Routes>
  );
}

export default function App() {
  return (
    <BrowserRouter>
      <ThemeProvider>
        <AuthProvider>
          <ToastProvider>
            <AppRoutes />
            <ToastContainer />
          </ToastProvider>
        </AuthProvider>
      </ThemeProvider>
    </BrowserRouter>
  );
}