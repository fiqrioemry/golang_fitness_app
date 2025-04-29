// src/pages/CreateNewProduct.jsx
import { Button } from "@/components/ui/button";
import { createProductSchema } from "@/lib/schema";
import { createProductState } from "@/lib/constant";
import { FormInput } from "@/components/form/FormInput";
import { useCreateProduct } from "@/hooks/useProductMutation";
import { SubmitButton } from "@/components/form/SubmitButton";
import { InputElement } from "@/components/input/InputElement";
import { useFieldArray, useFormContext } from "react-hook-form";
import { UploadElement } from "@/components/input/UploadElement";
import { SwitchElement } from "@/components/input/SwitchElement";
import { SelectCategoryElement } from "@/components/input/SelectCategoryElement";

const CreateNewProduct = () => {
  const { mutateAsync: createProduct } = useCreateProduct();

  return (
    <div className="max-w-2xl p-2">
      <h2 className="text-center">Create New Product</h2>
      <FormInput
        action={createProduct}
        state={createProductState}
        schema={createProductSchema}
      >
        {(methods) => (
          <>
            <InputElement
              name="name"
              label="Product Name"
              placeholder="Enter product name"
            />
            <InputElement
              rows={4}
              isTextArea
              name="description"
              label="Description"
              placeholder="Enter description"
            />
            <div>
              <SelectCategoryElement />
            </div>

            <SwitchElement name="isFeatured" label="Featured Product" />

            <div className="grid grid-cols-2 gap-4">
              <InputElement isNumber name="weight" label="Weight (g)" />
              <InputElement isNumber name="length" label="Length (cm)" />
              <InputElement isNumber name="width" label="Width (cm)" />
              <InputElement isNumber name="height" label="Height (cm)" />
            </div>

            <InputElement name="discount" label="Discount (%)" isNumber />

            <UploadElement name="images" label="Product Images" />

            <VariantSection />
            <AttributeSection />

            <SubmitButton text="Create Product" className="mt-6" />
          </>
        )}
      </FormInput>
    </div>
  );
};

export default CreateNewProduct;

// ================= VARIANT SECTION =================

const VariantSection = () => {
  const { control } = useFormContext();
  const { fields, append, remove } = useFieldArray({
    control,
    name: "variants",
  });

  return (
    <div className="space-y-4 ">
      <div className="flex justify-between items-center">
        <h4 className="font-semibold text-lg">Variants</h4>
        <Button
          type="button"
          variant="secondary"
          onClick={() =>
            append({
              sku: "",
              price: 0,
              stock: 0,
              sold: 0,
              isActive: true,
              imageUrl: "",
              options: {},
            })
          }
        >
          Add Variant
        </Button>
      </div>

      {fields.map((field, index) => (
        <div key={field.id} className="border p-4 rounded space-y-4">
          <div className="grid grid-cols-6 gap-2">
            <InputElement
              label="SKU"
              name={`variants.${index}.sku`}
              className="col-span-3"
              placeholder="Enter SKU"
            />

            <InputElement
              name={`variants.${index}.price`}
              label="Price"
              isNumber
            />
            <InputElement
              name={`variants.${index}.stock`}
              label="Stock"
              isNumber
            />
            <InputElement
              name={`variants.${index}.sold`}
              label="Sold"
              isNumber
            />
          </div>

          <SwitchElement name={`variants.${index}.isActive`} label="Active" />
          <div className="flex gap-2">
            <div className="w-1/2">
              <UploadElement
                name="imageUrl"
                label="Variant Product Image"
                isSingle
              />
            </div>
          </div>

          {/* Options - dynamic key value */}
          <div className="grid grid-cols-2 gap-4">
            <InputElement
              name={`variants.${index}.options.type`}
              label="Option Type (e.g Size, Color)"
            />
            <InputElement
              name={`variants.${index}.options.value`}
              label="Option Value (e.g M, Red)"
            />
          </div>

          <Button type="button" variant="danger" onClick={() => remove(index)}>
            Remove Variant
          </Button>
        </div>
      ))}
    </div>
  );
};

// ================= ATTRIBUTE SECTION =================

const AttributeSection = () => {
  const { control } = useFormContext();
  const { fields, append, remove } = useFieldArray({
    control,
    name: "attributes",
  });

  return (
    <div className="space-y-4 mt-6">
      <div className="flex justify-between items-center">
        <h4 className="font-semibold text-lg">Attributes</h4>
        <Button
          type="button"
          variant="secondary"
          onClick={() =>
            append({
              attributeId: 0,
              attributeValueId: 0,
            })
          }
        >
          Add Attribute
        </Button>
      </div>

      {fields.map((field, index) => (
        <div key={field.id} className="grid grid-cols-2 gap-4 items-end">
          <InputElement
            name={`attributes.${index}.attributeId`}
            label="Attribute ID"
            isNumber
          />
          <InputElement
            name={`attributes.${index}.attributeValueId`}
            label="Attribute Value ID"
            isNumber
          />

          <Button type="button" variant="danger" onClick={() => remove(index)}>
            Remove
          </Button>
        </div>
      ))}
    </div>
  );
};
