export const validateCardNumber = (value) => {
  const digits = value.replace(/\D/g, '');
  if (!digits) return { valid: false, error: 'Введите номер карты' };
  if (digits.length !== 16) return { valid: false, error: 'Номер карты должен содержать 16 цифр' };
  return { valid: true, value: digits };
};

export const validateAmount = (value) => {
  const num = Number(value);
  if (isNaN(num)) return { valid: false, error: 'Введите число' };
  if (num < 0) return { valid: false, error: 'Сумма не может быть отрицательной' };
  return { valid: true, value: num };
};

export const validateRequired = (value, fieldName = 'Поле') => {
  if (!value && value !== 0) {
    return { valid: false, error: `${fieldName} обязательно` };
  }
  return { valid: true, value };
};

export const validateEmail = (value) => {
  if (!value) return { valid: true, value };
  const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!regex.test(value)) {
    return { valid: false, error: 'Неверный формат email' };
  }
  return { valid: true, value };
};