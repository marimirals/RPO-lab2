import { useFormValidation, validators } from '../../hooks/useFormValidation';
import { formatCardNumber } from '../../utils/formatters';

export default function CardForm({ initialData = {}, onSubmit, onCancel, isEditing = false }) {
  const { values, errors, touched, handleChange, handleBlur, validateAll, reset } = useFormValidation(
    {
      card_number: initialData.card_number || '',
      balance: initialData.balance || 0,
      owner_name: initialData.owner_name || '',
      key_id: initialData.key_id || null,
      is_blocked: initialData.is_blocked || false,
    },
    {
      card_number: validators.cardNumber,
      balance: validators.positiveNumber,
      owner_name: validators.required('Введите имя владельца'),
    }
  );

  const handleSubmit = (e) => {
    e.preventDefault();
    if (validateAll()) {
      onSubmit({
        ...values,
        card_number: values.card_number.replace(/\D/g, ''),
      });
    }
  };

  const handleCardNumberChange = (e) => {
    const raw = e.target.value.replace(/\D/g, '').slice(0, 16);
    const formatted = formatCardNumber(raw);
    
    const syntheticEvent = {
      target: { name: 'card_number', value: formatted, type: 'text' }
    };
    handleChange(syntheticEvent);
  };

  return (
    <form onSubmit={handleSubmit} className="form">
      <div className="form-group">
        <label>Номер карты *</label>
        <input
          type="text"
          name="card_number"
          value={values.card_number}
          onChange={handleCardNumberChange}
          onBlur={handleBlur}
          placeholder="1234 5678 9012 3456"
          className={`input-with-mask ${touched.card_number && errors.card_number ? 'error' : ''}`}
        />
        {touched.card_number && errors.card_number && (
          <span className="form-error">{errors.card_number}</span>
        )}
        <span className="form-hint">16 цифр без пробелов</span>
      </div>

      <div className="form-group">
        <label>Баланс (копейки)</label>
        <input
          type="number"
          name="balance"
          value={values.balance}
          onChange={handleChange}
          onBlur={handleBlur}
          placeholder="0"
          min="0"
        />
      </div>

      <div className="form-group">
        <label>Имя владельца *</label>
        <input
          type="text"
          name="owner_name"
          value={values.owner_name}
          onChange={handleChange}
          onBlur={handleBlur}
          placeholder="Иван Иванов"
        />
        {touched.owner_name && errors.owner_name && (
          <span className="form-error">{errors.owner_name}</span>
        )}
      </div>

      <div className="form-checkbox">
        <input
          type="checkbox"
          name="is_blocked"
          checked={values.is_blocked}
          onChange={handleChange}
          id="is_blocked"
        />
        <label htmlFor="is_blocked">Заблокирована</label>
      </div>

      <div className="form-actions">
        <button type="submit">{isEditing ? 'Сохранить' : 'Создать'}</button>
        {onCancel && (
          <button type="button" className="secondary" onClick={onCancel}>
            Отмена
          </button>
        )}
      </div>
    </form>
  );
}