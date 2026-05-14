export default function Card({ children, className = '', hover = true }) {
  return (
    <div className={`card ${hover ? 'card--hover' : ''} ${className}`}>
      {children}
    </div>
  );
}

Card.Header = function CardHeader({ children, title, subtitle, actions }) {
  return (
    <div className="card-header">
      <div>
        {title && <div className="card-title">{title}</div>}
        {subtitle && <div className="card-subtitle">{subtitle}</div>}
      </div>
      {actions && <div className="card-actions">{actions}</div>}
    </div>
  );
};

Card.Body = function CardBody({ children }) {
  return <div className="card-body">{children}</div>;
};

Card.Row = function CardRow({ label, value, children }) {
  if (children) return <div className="card-row">{children}</div>;
  
  return (
    <div className="card-row">
      {label && <label>{label}</label>}
      {value && <value>{value}</value>}
    </div>
  );
};

Card.Footer = function CardFooter({ children }) {
  return <div className="card-footer">{children}</div>;
};