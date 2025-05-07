import { z } from "zod";

const imageItemSchema = z
  .any()
  .optional()
  .refine(
    (file) =>
      !file || // kosong = lolos
      file instanceof File ||
      typeof file === "string",
    { message: "Input must be a file or a valid URL" }
  )
  .refine(
    (file) =>
      !file || typeof file === "string" || file.type?.startsWith("image/"),
    { message: "File must be an image" }
  )
  .refine(
    (file) => !file || typeof file === "string" || file.size <= 2 * 1024 * 1024,
    { message: "Image size must be <= 2MB" }
  );

const requiredNumberField = (label, { min, max } = {}) =>
  z.preprocess(
    (val) => (val === "" || val === null ? undefined : val),
    z
      .number({
        required_error: `${label} is required`,
        invalid_type_error: `${label} must be a number`,
      })
      .refine(
        (val) =>
          (min === undefined || val >= min) &&
          (max === undefined || val <= max),
        {
          message:
            min !== undefined && max !== undefined
              ? `${label} must be between ${min} and ${max}`
              : min !== undefined
              ? `${label} must be at least ${min}`
              : max !== undefined
              ? `${label} must be at most ${max}`
              : `${label} must be valid`,
        }
      )
  );

// Register
export const sendOTPSchema = z.object({
  email: z.string().email("Invalid email address"),
});

export const verifyOTPSchema = z.object({
  email: z.string().email("Invalid email address"),
  otp: z.string().min(6, "OTP code must be at least 6 characters"),
});

export const registerSchema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string().min(6, "Password must be at least 6 characters"),
  fullname: z.string().min(1, "Full name is required"),
});

// Login
export const loginSchema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string().min(6, "Password must be at least 6 characters"),
  rememberMe: z.boolean().optional(),
});

// profiles
export const profileSchema = z.object({
  fullname: z.string().min(6, "Fullname must be at least 6 characters"),
  birthday: z.string().refine((val) => !isNaN(Date.parse(val)), {
    message: "Tanggal tidak valid",
  }),
  gender: z.string().optional(),
  phone: z.string().optional(),
  bio: z.string().optional(),
});

export const avatarSchema = z.object({
  avatar: imageItemSchema.refine((val) => !!val, {
    message: "Image is required",
  }),
});

// package
export const packageSchema = z.object({
  name: z.string().min(6, "Package name must be at least 6 characters"),
  description: z.string().min(20, "Description must be at least 20 characters"),
  price: z.number().positive("Price must be greater than 0"),
  credit: z.number().positive("Credit must be greater than 0"),
  discount: z
    .number()
    .min(0, "Discount cannot be negative")
    .max(100, "Discount cannot exceed 100")
    .optional(),
  expired: z.number().min(1, "Expiry duration is required"),
  additional: z.array(z.string()).optional(),
  image: imageItemSchema.refine((val) => !!val, {
    message: "Image is required",
  }),
  isActive: z.boolean().optional(),
});

// options
export const optionSchema = z.object({
  name: z.string().min(2, "Name must be at least 2 characters"),
});

export const subcategorySchema = z.object({
  name: z.string().min(2, "Name must be at least 2 characters"),
  categoryId: z.string().min(1, "Category is required"),
});

export const locationSchema = z.object({
  name: z.string().min(2, "Location name is required"),
  address: z.string().min(1, "Address is required"),
  geoLocation: z.string().min(1, "Geolocation is required"),
});

export const instructorSchema = z.object({
  userId: z.string().min(1, "User is required"),
  experience: z.number().optional(),
  specialties: z.string().optional(),
  certifications: z.string().optional(),
});

export const midtransNotificationSchema = z.object({
  transaction_status: z.string(),
  order_id: z.string(),
  payment_type: z.string(),
  fraud_status: z.string(),
});

export const createPaymentSchema = z.object({
  packageId: z.string().min(1, "Package is required"),
});

export const markAttendanceSchema = z.object({
  bookingId: z.string().min(1, "Booking ID is required"),
  status: z.enum(["attended", "absent", "cancelled"]),
});
export const reviewSchema = z.object({
  classId: z.string().min(1, "Class is required"),
  rating: z.number().min(1).max(5),
  comment: z.string().optional(),
});

export const bookingSchema = z.object({
  classScheduleId: z.string().min(1, "Class Schedule is required"),
});

export const scheduleSchema = z.object({
  isRecurring: z.boolean().optional().default(false),
  classId: z.string().min(1, "Class is required"),
  instructorId: z.string().min(1, "Instructor is required"),
  capacity: z.number().positive(),
  color: z.string().optional(),

  // Non-recurring
  date: z
    .string()
    .optional()
    .refine((val) => !val || !isNaN(Date.parse(val)), {
      message: "Date must be a valid date",
    }),

  // Time
  startHour: z.number().min(0).max(23),
  startMinute: z.number().max(59),

  // Recurring
  recurringDays: z
    .array(z.number().int().min(0).max(6), {
      required_error: "Recurring days must be an array of valid weekdays (0-6)",
    })
    .optional(),
  endType: z.enum(["never", "until"]).optional(),
  endDate: z
    .string()
    .optional()
    .refine((val) => !val || !isNaN(Date.parse(val)), {
      message: "End Date must be a valid date",
    }),
  capacity: requiredNumberField("Capacity", { min: 1 }),
});

// classes
export const classSchema = z.object({
  title: z.string().min(6, "Title must be at least 6 characters"),
  duration: requiredNumberField("Duration", { min: 15, max: 180 }),
  description: z.string().min(1, "Description is required"),
  additional: z.array(z.string()).optional(),
  typeId: z.string().min(1, "Type is required"),
  levelId: z.string().min(1, "Level is required"),
  locationId: z.string().min(1, "Location is required"),
  categoryId: z.string().min(1, "Category is required"),
  subcategoryId: z.string().min(1, "Subcategory is required"),
  image: imageItemSchema.refine((val) => !!val, {
    message: "Image is required",
  }),
  isActive: z.boolean().optional(),
});

export const uploadGallerySchema = z.object({
  images: z.array(imageItemSchema).min(1, "Image is required"),
});
