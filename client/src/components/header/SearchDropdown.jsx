const SearchDropdown = ({ search, isLoading, results, onClick }) => {
  if (!search) return null;
  console.log(results);
  return (
    <div className="absolute top-full left-0 right-0 bg-white border mt-1 rounded-md shadow-md max-h-80 overflow-y-auto z-50">
      {isLoading ? (
        <div className="p-3 text-sm text-gray-500">Loading...</div>
      ) : results?.length ? (
        results.map((item) => (
          <div
            key={item.ID}
            onClick={() => onClick(item.slug)}
            className="flex items-center gap-3 p-2 hover:bg-gray-100 cursor-pointer"
          >
            <img
              src={item.images[0]}
              alt={item.name}
              className="w-10 h-10 object-cover rounded"
            />
            <span className="text-sm">{item.name}</span>
          </div>
        ))
      ) : (
        <div className="p-2 text-sm text-gray-500">No result</div>
      )}
    </div>
  );
};

export default SearchDropdown;
