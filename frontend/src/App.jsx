import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import Header from './components/Layout/Header';
import Sidebar from './components/Layout/Sidebar';
import ProtectedRoute from './components/Layout/ProtectedRoute';

import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import Cards from './pages/Cards';
import Terminals from './pages/Terminals';
import Transactions from './pages/Transactions';
import Keys from './pages/Keys';
import TerminalTest from './pages/TerminalTest';

function Layout({ children }) {
  return (
    <div style={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      <Header />
      <div style={{ display: 'flex', flex: 1 }}>
        <Sidebar />
        <main style={{ flex: 1, padding: '1rem' }}>{children}</main>
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
      <AuthProvider>
        <AppRoutes />
      </AuthProvider>
    </BrowserRouter>
  );
}