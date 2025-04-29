const FetchLoading = () => {
  return (
    <div className="flex items-center justify-center h-[50vh]">
      <div className="animate-spin rounded-full h-12 w-12 border-4 border-t-primary border-gray-200" />
      <span className="ml-4 text-sm text-muted-foreground animate-pulse">
        Memuat produk...
      </span>
    </div>
  );
};

export default FetchLoading;
