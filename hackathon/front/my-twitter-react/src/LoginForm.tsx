// LoginForm.tsx
import React, { useState } from 'react';
import { fireAuth, signInWithEmailAndPassword } from './firebase';

const LoginForm: React.FC = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState<string | null>(null); // エラーステートを追加

    const handleLogin = async () => {
        try {
            await signInWithEmailAndPassword(fireAuth, email, password);
            console.log('User logged in successfully');
            // ログイン成功後の処理を追加する場合はここに記述
        } catch (error: any) {
            console.error('Error logging in:', error.message);
            setError('Error: Authentication failed. Please reenter this form.');
        }
    };

    return (
        <div>
            <h2>Login</h2>
            <p>for registered user</p>
            {error && <p style={{ color: 'red' }}>{error}</p>} {/* エラーメッセージを表示 */}
            <input type="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
            <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
            <button onClick={handleLogin}>Login</button>
        </div>
    );
};

export default LoginForm;
