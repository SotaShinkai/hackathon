// LogoutForm.tsx
import React, { useState } from 'react';
import { fireAuth, signOut} from './firebase';

const LogoutForm: React.FC = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState<string | null>(null); // エラーステートを追加

    const handleLogout = async () => {
        try {
            await signOut(fireAuth);
            console.log('User logged out successfully');
            // ログイン成功後の処理を追加する場合はここに記述
        } catch (error: any) {
            console.error('Error logged out:', error.message);
            setError('Error: Authentication failed. Please reenter this form.');
        }
    };

    return (
        <div>
            <button onClick={handleLogout}>Logout</button>
        </div>
    );
};

export default LogoutForm;
