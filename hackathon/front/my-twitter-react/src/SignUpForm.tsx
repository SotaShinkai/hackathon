import React, { useState } from 'react';
import { fireAuth, createUserWithEmailAndPassword } from './firebase';

const SignUpForm: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);

  const handleSignUp = async () => {
    try {
      await createUserWithEmailAndPassword(fireAuth, email, password);
      console.log('User signed up successfully');
      // ユーザー登録成功後の処理を追加する場合はここに記述
    } catch (error: any) {
      console.error('Error signing up:', error.message);
      setError('Error: SignUp failed. Please reenter this form.');
    }
  };

  return (
      <div>
          <h2>Sign Up</h2>
          <p>for unregistered user</p>
          {error && <p style={{ color: 'red' }}>{error}</p>}
          <input type="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
          <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
          <button onClick={handleSignUp}>Sign Up</button>
      </div>
  );
};

export default SignUpForm;

