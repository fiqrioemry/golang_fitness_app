import { X } from "lucide-react";
import { useRef } from "react";

const SearchInput = ({ value, onChange, onKeyDown, onClear }) => {
  const inputRef = useRef(null);

  return (
    <div className="relative w-full">
      <input
        ref={inputRef}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        onKeyDown={onKeyDown}
        className="w-full border px-3 py-2 pr-10 rounded-md text-sm"
        placeholder="Search products..."
      />
      {value && (
        <button
          onClick={() => {
            onClear();
            inputRef.current?.focus();
          }}
          className="absolute right-2 top-1/2 -translate-y-1/2 text-gray-500 hover:text-red-500"
        >
          <X />
        </button>
      )}
    </div>
  );
};

export default SearchInput;
