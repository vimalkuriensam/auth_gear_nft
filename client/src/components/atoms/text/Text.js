import React from "react";

const Text = ({
  variant,
  children,
  style,
  className,
  edit = false,
  onHandleEdit = () => {},
}) => {
  if (edit) {
    return (
      <div
        className={`text text--${variant} ${className}`}
        contentEditable={true}
        onBlur={onHandleEdit}
        dangerouslySetInnerHTML={{ __html: children }}
        style={{ ...style }}
      />
    );
  }
  return (
    <div className={`text text--${variant} ${className}`} style={{ ...style }}>
      {children}
    </div>
  );
};

export default Text;
