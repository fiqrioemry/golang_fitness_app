import React from "react";
import { createVoucherSchema } from "@/lib/schema";
import { useVoucherMutation } from "@/hooks/useVouchers";
import { SwitchElement } from "@/components/input/SwitchElement";
import { SelectElement } from "@/components/input/SelectElement";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputDateElement } from "@/components/input/InputDateElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const discountTypeOptions = [
  { label: "Fixed", value: "fixed" },
  { label: "Percentage", value: "percentage" },
];

const VoucherUpdate = ({ voucher }) => {
  const { mutateAsync, isPending } = useVoucherMutation();

  return (
    <FormUpdateDialog
      state={voucher}
      loading={isPending}
      action={({ id, ...data }) => mutateAsync({ id, data })}
      title="Update Voucher"
      schema={createVoucherSchema}
    >
      <InputTextElement
        name="code"
        label="Voucher Code (cannot changed)"
        disabled
      />
      <InputTextareaElement
        name="description"
        label="Description"
        placeholder="e.g. Diskon 50% untuk semua kelas"
        maxLength={200}
      />
      <SelectElement
        name="discountType"
        label="Discount Type"
        options={discountTypeOptions}
        placeholder="Select discount type"
      />
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <InputNumberElement
          name="discount"
          label="Discount"
          placeholder="e.g. 50000 or 50"
        />
        <InputNumberElement
          name="maxDiscount"
          label="Max Discount (if %)"
          placeholder="e.g. 30000"
          isOptional
        />
        <InputNumberElement name="quota" label="Quota" placeholder="e.g. 10" />
      </div>
      <InputDateElement
        mode="future"
        name="expiredAt"
        label="Expiration Date"
        placeholder="YYYY-MM-DD"
      />
      <SwitchElement name="isReusable" label="Allow multiple usage?" />
    </FormUpdateDialog>
  );
};

export { VoucherUpdate };
