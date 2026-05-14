import { useState, useCallback } from 'react';

export const useFormValidation = (initialValues, validators) => {
  const [values, setValues] = useState(initialValues);
  const [errors, setErrors] = useState({});
  const [touched, setTouched] = useState({});

  const validateField = useCallback((name, value) => {
    const validator = validators[name];
    if (!validator) return '';
    
    const error = validator(value, values);
    return error || '';
  }, [validators, values]);

  const validateAll = useCallback(() => {
    const newErrors = {};
    let isValid = true;
    
    Object.keys(validators).forEach(name => {
      const error = validateField(name, values[name]);
      if (error) {
        newErrors[name] = error;
        isValid = false;
      }
    });
    
    setErrors(newErrors);
    return isValid;
  }, [values, validators, validateField]);

  const handleChange = useCallback((e) => {
    const { name, value, type, checked } = e.target;
    const newValue = type === 'checkbox' ? checked : value;
    
    setValues(prev => ({ ...prev, [name]: newValue }));
    
    if (touched[name]) {
      const error = validateField(name, newValue);
      setErrors(prev => ({ ...prev, [name]: error }));
    }
  }, [touched, validateField]);

  const handleBlur = useCallback((e) => {
    const { name, value } = e.target;
    setTouched(prev => ({ ...prev, [name]: true }));
    
    const error = validateField(name, value);
    setErrors(prev => ({ ...prev, [name]: error }));
  }, [validateField]);

  const reset = useCallback(() => {
    setValues(initialValues);
    setErrors({});
    setTouched({});
  }, [initialValues]);

  return {
    values,
    errors,
    touched,
    handleChange,
    handleBlur,
    validateAll,
    reset,
    setValues,
  };
};

// Validators
export const validators = {
  required: (value, message = 'Обязательное поле') => 
    !value && value !== 0 ? message : '',
  
  minLength: (min, message) => (value) =>
    value && value.length < min ? (message || `Минимум ${min} символов`) : '',
  
  maxLength: (max, message) => (value) =>
    value && value.length > max ? (message || `Максимум ${max} символов`) : '',
  
  pattern: (regex, message) => (value) =>
    value && !regex.test(value) ? (message || 'Неверный формат') : '',
  
  cardNumber: (value) => {
    const digits = value.replace(/\D/g, '');
    if (!digits) return 'Введите номер карты';
    if (digits.length !== 16) return 'Номер карты должен содержать 16 цифр';
    return '';
  },
  
  positiveNumber: (value, message = 'Значение должно быть положительным') =>
    value < 0 ? message : '',
};