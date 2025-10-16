// client/src/components/ProtectedRoute.tsx
import React, { useEffect, useState } from 'react';
import { Navigate } from 'react-router-dom';
import axios from 'axios';

type Props = {
  children?: React.ReactNode;
};

const API = import.meta.env.VITE_API_URL || '/api/auth'; // use relative path if using Vite proxy

export default function ProtectedRoute({ children }: Props) {
  const [loading, setLoading] = useState(true);
  const [authed, setAuthed] = useState(false);

  useEffect(() => {
    let mounted = true;
    async function validate() {
      try {
        // call a lightweight validate endpoint that uses your auth middleware
        const res = await axios.get(`${API}/validate`, { withCredentials: true });
        if (!mounted) return;
        if (res.status === 200) {
          setAuthed(true);
        } else {
          setAuthed(false);
        }
      } catch (err) {
        setAuthed(false);
      } finally {
        if (mounted) setLoading(false);
      }
    }
    validate();
    return () => { mounted = false; };
  }, []);

  if (loading) {
    // render a spinner or null while we check
    return <div style={{padding:20}}>Checking authentication…</div>;
  }

  if (!authed) {
    // not authenticated — redirect to login
    return <Navigate to="/login" replace />;
  }

  // authenticated — render children (the protected page)
  return <>{children}</>;
}