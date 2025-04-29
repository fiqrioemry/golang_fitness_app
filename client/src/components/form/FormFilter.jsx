import React, { useState, useEffect, useRef } from "react";
import { useSearchParams } from "react-router-dom";

const FormFilter = ({ children }) => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [filters, setFilters] = useState(() =>
    Object.fromEntries([...searchParams])
  );
  const debounceTimer = useRef(null);

  const updateParams = (newFilters) => {
    if (debounceTimer.current) clearTimeout(debounceTimer.current);

    debounceTimer.current = setTimeout(() => {
      const params = new URLSearchParams();
      Object.entries(newFilters).forEach(([key, value]) => {
        if (value) params.set(key, value);
      });
      setSearchParams(params);
    }, 300); // debounce 300ms
  };

  useEffect(() => {
    updateParams(filters);
    return () => {
      if (debounceTimer.current) clearTimeout(debounceTimer.current);
    };
  }, [filters, setSearchParams]);

  const handleChange = (name, value) => {
    setFilters((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleReset = () => {
    setFilters({});
    setSearchParams({});
  };

  const enhancedChildren = React.Children.map(children, (child) => {
    if (React.isValidElement(child) && child.props && child.props.name) {
      return React.cloneElement(child, {
        onChange: (e) => handleChange(child.props.name, e.target.value),
        value: filters[child.props.name] || "",
      });
    }
    return child;
  });

  return (
    <div className="space-y-4">
      <div className="flex flex-wrap gap-4">{enhancedChildren}</div>
      <button
        type="button"
        onClick={handleReset}
        className="px-4 py-2 text-sm bg-gray-200 rounded hover:bg-gray-300"
      >
        Reset Filters
      </button>
    </div>
  );
};

export default FormFilter;
