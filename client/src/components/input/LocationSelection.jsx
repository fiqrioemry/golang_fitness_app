import {
  useProvincesQuery,
  useDistrictsByCityQuery,
  useCitiesByProvinceQuery,
  useSubdistrictsByDistrictQuery,
  usePostalCodesBySubdistrictQuery,
} from "@/hooks/useLocationQuery";
import { useFormContext, Controller } from "react-hook-form";
import { useEffect, useRef } from "react";

const LocationSelection = () => {
  const { control, watch } = useFormContext();

  const selectedProvinceId = watch("provinceId");
  const selectedCityId = watch("cityId");
  const selectedDistrictId = watch("districtId");
  const selectedSubdistrictId = watch("subdistrictId");

  const { data: provinces = [] } = useProvincesQuery();
  const { data: cities = [] } = useCitiesByProvinceQuery(selectedProvinceId);
  const { data: districts = [] } = useDistrictsByCityQuery(selectedCityId);
  const { data: subdistricts = [] } =
    useSubdistrictsByDistrictQuery(selectedDistrictId);
  const { data: postalCodes = [] } = usePostalCodesBySubdistrictQuery(
    selectedSubdistrictId
  );

  // Refs
  const cityRef = useRef(null);
  const districtRef = useRef(null);
  const subdistrictRef = useRef(null);
  const postalCodeRef = useRef(null);

  // ScrollArea container ref
  const scrollAreaContainer = useRef(null);

  useEffect(() => {
    // Cari parent ScrollArea
    scrollAreaContainer.current = document.querySelector(".scroll-area");
  }, []);

  const scrollToRef = (ref) => {
    if (ref.current && scrollAreaContainer.current) {
      scrollAreaContainer.current.scrollTo({
        top: ref.current.offsetTop - 100, // kasih jarak dikit biar enak
        behavior: "smooth",
      });
    }
  };

  useEffect(() => {
    if (selectedProvinceId) scrollToRef(cityRef);
  }, [selectedProvinceId]);

  useEffect(() => {
    if (selectedCityId) scrollToRef(districtRef);
  }, [selectedCityId]);

  useEffect(() => {
    if (selectedDistrictId) scrollToRef(subdistrictRef);
  }, [selectedDistrictId]);

  useEffect(() => {
    if (selectedSubdistrictId) scrollToRef(postalCodeRef);
  }, [selectedSubdistrictId]);

  const SelectField = ({
    name,
    label,
    options,
    optionLabelKey = "name",
    optionValueKey = "id",
    innerRef,
  }) => (
    <Controller
      control={control}
      name={name}
      rules={{ required: true }}
      render={({ field }) => (
        <div ref={innerRef}>
          <label className="block mb-1 font-medium">{label}</label>
          <select
            {...field}
            onChange={(e) => field.onChange(Number(e.target.value))}
            value={field.value || ""}
            className="w-full border p-2 rounded"
          >
            <option value="">Select {label}</option>
            {options.map((option) => (
              <option
                key={option[optionValueKey]}
                value={option[optionValueKey]}
              >
                {option[optionLabelKey]}
              </option>
            ))}
          </select>
        </div>
      )}
    />
  );

  return (
    <div className="space-y-4">
      {/* Province */}
      <SelectField name="provinceId" label="Province" options={provinces} />

      {/* City */}
      {selectedProvinceId && (
        <SelectField
          name="cityId"
          label="City"
          options={cities}
          innerRef={cityRef}
        />
      )}

      {/* District */}
      {selectedCityId && (
        <SelectField
          name="districtId"
          label="District"
          options={districts}
          innerRef={districtRef}
        />
      )}

      {/* Subdistrict */}
      {selectedDistrictId && (
        <SelectField
          name="subdistrictId"
          label="Subdistrict"
          options={subdistricts}
          innerRef={subdistrictRef}
        />
      )}

      {/* Postal Code */}
      {selectedSubdistrictId && (
        <SelectField
          name="postalCodeId"
          label="Postal Code"
          options={postalCodes}
          optionLabelKey="postalCode"
          innerRef={postalCodeRef}
        />
      )}
    </div>
  );
};

export default LocationSelection;
