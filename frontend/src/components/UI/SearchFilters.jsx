import { useState } from 'react';

export default function SearchFilters({ 
  searchPlaceholder = 'Поиск...',
  filters = [],
  onSearch,
  onFilterChange,
  children 
}) {
  const [searchQuery, setSearchQuery] = useState('');
  const [activeFilters, setActiveFilters] = useState({});

  const handleSearch = (e) => {
    e.preventDefault();
    onSearch?.(searchQuery, activeFilters);
  };

  const handleFilterToggle = (filterName, value) => {
    const newFilters = { ...activeFilters, [filterName]: value };
    setActiveFilters(newFilters);
    onFilterChange?.(newFilters);
  };

  const clearFilters = () => {
    setActiveFilters({});
    setSearchQuery('');
    onFilterChange?.({});
    onSearch?.('', {});
  };

  const hasActiveFilters = Object.values(activeFilters).some(v => v) || searchQuery;

  return (
    <div className="search-filters">
      <form onSubmit={handleSearch} className="form" style={{ flex: 1, display: 'flex', gap: '0.75rem' }}>
        <div className="form-group" style={{ flex: 1, margin: 0 }}>
          <input
            type="text"
            placeholder={searchPlaceholder}
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            style={{ margin: 0 }}
          />
        </div>
        <button type="submit">Найти</button>
        {hasActiveFilters && (
          <button type="button" className="secondary" onClick={clearFilters}>
            Сбросить
          </button>
        )}
      </form>
      
      {filters.length > 0 && (
        <div className="filter-chips">
          {filters.map(filter => (
            <div key={filter.name} className="filter-chip">
              {filter.label}: {activeFilters[filter.name] || filter.default || '—'}
              {activeFilters[filter.name] && (
                <button onClick={() => handleFilterToggle(filter.name, '')}>&times;</button>
              )}
            </div>
          ))}
        </div>
      )}
      
      {children}
    </div>
  );
}