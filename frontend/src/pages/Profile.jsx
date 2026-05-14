import { useState } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { useToast } from '../hooks/useToast';
import Card from '../components/UI/Card';

export default function Profile() {
  const { user } = useAuth();
  const { toast } = useToast();
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState({
    name: user?.name || '',
    login: user?.login || '',
  });

  const handleSubmit = async (e) => {
    e.preventDefault();
    // Здесь будет запрос на обновление профиля
    toast.success('Профиль обновлён');
    setIsEditing(false);
  };

  if (!user) return <div className="page">Загрузка...</div>;

  return (
    <div className="page">
      <h2>Мой профиль</h2>
      
      <div className="flex" style={{ gap: '2rem', marginTop: '1.5rem' }}>
        <Card style={{ flex: 1, maxWidth: '400px' }}>
          <Card.Header title="Личная информация" />
          <Card.Body>
            <Card.Row label="Логин" value={user.login} />
            <Card.Row label="Имя" value={user.name} />
            <Card.Row label="Роль" value={user.is_admin ? 'Администратор' : 'Пользователь'} />
            <Card.Row label="Дата регистрации" value={new Date(user.created_at).toLocaleDateString('ru-RU')} />
          </Card.Body>
        </Card>

        <Card style={{ flex: 2 }}>
          <Card.Header 
            title={isEditing ? 'Редактирование' : 'Мои карты'} 
            actions={
              !isEditing && (
                <button className="secondary" onClick={() => setIsEditing(true)}>
                  Редактировать
                </button>
              )
            }
          />
          <Card.Body>
            {isEditing ? (
              <form onSubmit={handleSubmit} className="form">
                <div className="form-group">
                  <label>Имя</label>
                  <input
                    type="text"
                    value={formData.name}
                    onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                  />
                </div>
                <div className="form-actions">
                  <button type="submit">Сохранить</button>
                  <button type="button" className="secondary" onClick={() => setIsEditing(false)}>
                    Отмена
                  </button>
                </div>
              </form>
            ) : (
              <div>
                <p className="text-muted">Здесь будет список ваших карт</p>
              </div>
            )}
          </Card.Body>
        </Card>
      </div>
    </div>
  );
}