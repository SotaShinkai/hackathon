// App.tsx
import React, { useState, useEffect } from 'react';
import SignUpForm from './SignUpForm';
import LoginForm from './LoginForm';
import Contents from './Contents';
import ContentsFail from './ContentsFail'
import { fireAuth } from './firebase'; // firebaseからfireAuthをインポート
import { User, onAuthStateChanged } from 'firebase/auth'; // 必要なFirebase Authの型をインポート

const App: React.FC = () => {
  const [loginUser, setLoginUser] = useState<User | null>(null); // User型を明示的に指定

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(fireAuth, (user) => {
      setLoginUser(user);
    });

    return () => unsubscribe();
  }, []);

  return (
      <div>
        <h1>Authentication</h1>
        {(
            <>
              <SignUpForm />
              <LoginForm />
            </>
        )}
        {loginUser ? <Contents /> : <ContentsFail />}
      </div>
  );
};

export default App;

