// ProductResults.jsx with pagination integration
import {
  Select,
  SelectGroup,
  SelectItem,
  SelectValue,
  SelectContent,
  SelectTrigger,
} from "@/components/ui/select";

import { useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { Grid2X2, Grid3X3, X } from "lucide-react";
import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import { useCategoriesQuery } from "@/hooks/useCategoryQuery";
import { useSearchProductsQuery } from "@/hooks/useProductsQuery";
import ProductCard from "@/components/product-results/ProductCard";
import ProductList from "@/components/product-results/ProductList";
import NoProductResult from "@/components/product-results/NoProductResult";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

const ProductResults = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [showPriceWarning, setShowPriceWarning] = useState(false);

  const queryParams = useMemo(() => {
    const params = Object.fromEntries(searchParams.entries());
    return {
      ...params,
      minPrice: parseFloat(params.minPrice) || undefined,
      maxPrice: parseFloat(params.maxPrice) || undefined,
      page: parseInt(params.page || "1"),
    };
  }, [searchParams]);

  const {
    data: searchData,
    isLoading,
    isError,
    refetch,
  } = useSearchProductsQuery(queryParams);

  const { data: categoriesData, isLoading: isCategoriesLoading } =
    useCategoriesQuery();

  const categories = categoriesData?.categories || [];
  const [minPriceInput, setMinPriceInput] = useState(
    queryParams.minPrice || ""
  );
  const [maxPriceInput, setMaxPriceInput] = useState(
    queryParams.maxPrice || ""
  );

  const handleFilterChange = (key, value) => {
    const params = new URLSearchParams(searchParams);
    if (value) params.set(key, value);
    else params.delete(key);
    if (params.get("page")) params.delete("page");
    setSearchParams(params);
  };

  const handlePageChange = (newPage) => {
    const params = new URLSearchParams(searchParams);
    if (newPage > 1) {
      params.set("page", newPage);
    } else {
      params.delete("page");
    }
    setSearchParams(params);
  };

  const handlePriceInputChange = (key, value) => {
    if (!/^[0-9]*$/.test(value)) return;
    key === "minPrice" ? setMinPriceInput(value) : setMaxPriceInput(value);
  };

  const applyPriceFilter = () => {
    const minVal = parseFloat(minPriceInput);
    const maxVal = parseFloat(maxPriceInput);

    setShowPriceWarning(minVal && maxVal && minVal > maxVal);

    const params = new URLSearchParams(searchParams);
    if (minPriceInput) params.set("minPrice", minPriceInput);
    else params.delete("minPrice");
    if (maxPriceInput) params.set("maxPrice", maxPriceInput);
    else params.delete("maxPrice");
    if (params.get("page")) params.delete("page");
    setSearchParams(params);
  };

  const activeFilters = Array.from(searchParams.entries()).filter(([key]) =>
    ["q", "category", "subcategory", "minPrice", "maxPrice"].includes(key)
  );

  const removeFilter = (key) => {
    const params = new URLSearchParams(searchParams);
    params.delete(key);
    if (params.get("page")) params.delete("page");
    setSearchParams(params);
  };

  const clearAllFilters = () => {
    setSearchParams({});
  };

  if (isLoading || isCategoriesLoading) return <FetchLoading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  const { results = [], pages = 1, page = 1 } = searchData || {};

  return (
    <section className="section">
      {showPriceWarning && (
        <div className="bg-red-100 text-red-600 text-sm p-3 rounded border border-red-300">
          Harga maksimum harus lebih besar dari harga minimum.
        </div>
      )}

      <div className="grid grid-cols-4 gap-4">
        <div className="col-span-4 md:col-span-1 space-y-4">
          {/* Category Filter with Checkbox and Scrollable */}
          <div className="space-y-2">
            <h3 className="text-lg font-semibold text-gray-700">Kategori</h3>
            <div className="h-64 overflow-y-auto border rounded-lg p-3 shadow-sm bg-white">
              {categories.map((cat) => (
                <div key={cat.ID} className="mb-3">
                  <label className="flex items-center gap-2 font-medium cursor-pointer">
                    <input
                      type="checkbox"
                      checked={searchParams.get("category") === cat.slug}
                      onChange={() =>
                        handleFilterChange(
                          "category",
                          cat.slug === searchParams.get("category")
                            ? ""
                            : cat.slug
                        )
                      }
                      className="accent-primary"
                    />
                    {cat.name}
                  </label>
                  {cat.Subcategories?.length > 0 && (
                    <div className="ml-6 mt-1 space-y-1">
                      {cat.Subcategories.map((sub) => (
                        <label
                          key={sub.slug}
                          className="flex items-center gap-2 text-sm text-muted-foreground cursor-pointer"
                        >
                          <input
                            type="checkbox"
                            checked={
                              searchParams.get("subcategory") === sub.slug
                            }
                            onChange={() =>
                              handleFilterChange(
                                "subcategory",
                                sub.slug === searchParams.get("subcategory")
                                  ? ""
                                  : sub.slug
                              )
                            }
                            className="accent-primary"
                          />
                          {sub.name}
                        </label>
                      ))}
                    </div>
                  )}
                </div>
              ))}
            </div>
          </div>

          {/* Price Filter */}
          <div className="space-y-3 pt-6 border-t">
            <h3 className="text-lg font-semibold text-gray-700">Harga</h3>
            <div className="flex gap-2">
              <input
                type="text"
                inputMode="numeric"
                value={minPriceInput}
                onChange={(e) =>
                  handlePriceInputChange("minPrice", e.target.value)
                }
                className="border rounded px-3 py-2 w-full text-sm focus:outline-none focus:ring-1 focus:ring-primary"
                placeholder="Min"
              />
              <input
                type="text"
                inputMode="numeric"
                value={maxPriceInput}
                onChange={(e) =>
                  handlePriceInputChange("maxPrice", e.target.value)
                }
                className="border rounded px-3 py-2 w-full text-sm focus:outline-none focus:ring-1 focus:ring-primary"
                placeholder="Max"
              />
            </div>
            <button
              onClick={applyPriceFilter}
              className="bg-primary hover:bg-primary/90 text-white text-sm rounded-md px-4 py-2 w-full transition"
            >
              Terapkan Filter
            </button>
          </div>
        </div>

        <div className="col-span-4 md:col-span-3">
          <div className="mb-4">
            <Tabs defaultValue="view1">
              <TabsList className="flex items-center justify-between mb-4">
                <div>
                  <TabsTrigger value="view1">
                    <Grid2X2 />
                  </TabsTrigger>
                  <TabsTrigger value="view2">
                    <Grid3X3 />
                  </TabsTrigger>
                </div>
                <div>
                  <Select
                    onValueChange={(val) => handleFilterChange("sort", val)}
                  >
                    <SelectTrigger className="w-[180px]">
                      <SelectValue placeholder="Sort by" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        <SelectItem value="price:asc">Lower Price</SelectItem>
                        <SelectItem value="price:desc">Higher Price</SelectItem>
                        <SelectItem value="createdAt:asc">Newest</SelectItem>
                        <SelectItem value="createdAt:desc">Oldest</SelectItem>
                      </SelectGroup>
                    </SelectContent>
                  </Select>
                </div>
              </TabsList>

              {activeFilters.length > 0 && (
                <div className="flex flex-wrap items-center gap-2 mb-2 py-2 border-b">
                  {activeFilters.map(([key, val]) => (
                    <span
                      key={key}
                      className="bg-muted text-sm px-3 py-1 rounded-full flex items-center gap-1"
                    >
                      {key === "minPrice"
                        ? "Harga Minimum"
                        : key === "maxPrice"
                        ? "Harga Maksimum"
                        : val}
                      <button onClick={() => removeFilter(key)}>
                        <X size={14} />
                      </button>
                    </span>
                  ))}
                  <button
                    onClick={clearAllFilters}
                    className="text-sm text-green-600 hover:underline"
                  >
                    Hapus Semua
                  </button>
                </div>
              )}

              {results.length === 0 ? (
                <NoProductResult />
              ) : (
                <>
                  <TabsContent
                    className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6"
                    value="view1"
                  >
                    {results.map((product) => (
                      <ProductCard product={product} />
                    ))}
                  </TabsContent>
                  <TabsContent className="space-y-4" value="view2">
                    {results.map((product) => (
                      <ProductList key={product.id} product={product} />
                    ))}
                  </TabsContent>

                  <div className="mt-6 flex justify-center gap-2 text-sm">
                    {Array.from({ length: pages }, (_, i) => (
                      <button
                        key={i + 1}
                        onClick={() => handlePageChange(i + 1)}
                        className={`border px-3 py-1 rounded ${
                          i + 1 === page ? "bg-primary text-white" : ""
                        }`}
                      >
                        {i + 1}
                      </button>
                    ))}
                  </div>
                </>
              )}
            </Tabs>
          </div>
        </div>
      </div>
    </section>
  );
};

export default ProductResults;
