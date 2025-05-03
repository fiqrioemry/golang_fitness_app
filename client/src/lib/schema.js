import { z } from "zod";

const imageItemSchema = z.union([
  z
    .instanceof(File)
    .refine((file) => file.type.startsWith("image/"), {
      message: "File must be an image",
    })
    .refine((file) => file.size <= 2 * 1024 * 1024, {
      message: "Image size must be less than or equal to 2MB",
    }),
  z.string().url(),
]);

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
});

// profiles
export const profileSchema = z.object({
  fullname: z.string().min(6, "Full name must be at least 6 characters"),
  birthday: z.string(),
  gender: z.string(),
  phone: z.string(),
  bio: z.string(),
});

export const avatarSchema = z.object({
  avatar: z.union([z.instanceof(File), z.string().url()]),
});

// classes
export const classSchema = z.object({
  title: z.string().min(6, "Title must be at least 6 characters"),
  duration: z
    .number()
    .min(15, "Minimum duration is 15 minutes")
    .max(180, "Maximum duration is 180 minutes"),
  description: z.string().min(1, "Description is required"),
  additional: z.array(z.string()).optional(),
  typeId: z.string().min(1, "Type is required"),
  levelId: z.string().min(1, "Level is required"),
  locationId: z.string().min(1, "Location is required"),
  categoryId: z.string().min(1, "Category is required"),
  subcategoryId: z.string().min(1, "Subcategory is required"),
  image: z.union([z.instanceof(File), z.string().url()]),
  images: z.array(imageItemSchema).optional(),
  isActive: z.boolean().optional(),
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
    .max(100, "Discount cannot exceed 100"),
  isActive: z.boolean(),
  expired: z.number().min(1, "Expiry duration is required"),
  additional: z.array(z.string()).optional(),
  image: z.union([z.instanceof(File), z.string().url()]),
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

export const createScheduleTemplateSchema = z.object({
  classId: z.string().min(1, "Class is required"),
  instructorId: z.string().min(1, "Instructor is required"),
  dayOfWeek: z.number().min(0).max(6),
  startHour: z.number().min(0).max(23),
  startMinute: z.number().min(0).max(59),
  capacity: z.number().positive(),
});

export const updateClassScheduleSchema = z.object({
  startTime: z.string().optional(),
  endTime: z.string().optional(),
  capacity: z.number().optional(),
});

export const createClassScheduleSchema = z.object({
  classId: z.string().min(1, "Class is required"),
  instructorId: z.string().min(1, "Instructor is required"),
  startTime: z.string().min(1, "Start time is required"),
  capacity: z.number().positive(),
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

export const classScheduleSchema = z.object({
  classId: z.string().min(1, "Class is required"),
  date: z.string().min(1, "Date is required"),
  instructorId: z.string().min(1, "Instructor is required"),
  startHour: z.number().min(8).max(17),
  startMinute: z.number().min(0).max(45),
  capacity: z.number().min(0, "Capacity must be more than 0"),
  isActive: z.boolean().optional(),
});
