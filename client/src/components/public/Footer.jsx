import {
  FaGithub,
  FaLinkedin,
  FaTwitter,
  FaInstagram,
  FaEnvelope,
} from "react-icons/fa";

const Footer = () => {
  return (
    <footer className="bg-background border-t text-muted-foreground">
      <div className=" max-w-7xl mx-auto px-4 py-10 grid gap-8 grid-cols-1 md:grid-cols-4">
        {/* Brand & Slogan */}
        <div>
          <h2 className="text-xl font-bold text-primary">FitBook Studio</h2>
          <p className="text-sm mt-2">
            Your go-to place for high-energy fitness classes and flexible
            membership packages.
          </p>
        </div>

        {/* Navigation Links */}
        <div>
          <h3 className="text-base font-semibold mb-2">Quick Links</h3>
          <ul className="space-y-1 text-sm">
            <li>
              <a href="/" className="hover:text-primary">
                Home
              </a>
            </li>
            <li>
              <a href="/classes" className="hover:text-primary">
                Classes
              </a>
            </li>
            <li>
              <a href="/packages" className="hover:text-primary">
                Packages
              </a>
            </li>
            <li>
              <a href="/schedule" className="hover:text-primary">
                Schedule
              </a>
            </li>
            <li>
              <a href="/contact" className="hover:text-primary">
                Contact
              </a>
            </li>
          </ul>
        </div>

        {/* Newsletter */}
        <div>
          <h3 className="text-base font-semibold mb-2">Join Our Community</h3>
          <p className="text-sm mb-3">
            Get tips, promotions, and new class updates straight to your inbox.
          </p>
          <form className="flex flex-col gap-2">
            <input
              type="email"
              placeholder="Enter your email"
              className="px-3 py-2 rounded border focus:outline-none focus:ring focus:border-primary text-sm"
            />
            <button
              type="submit"
              className="bg-primary text-white text-sm px-4 py-2 rounded hover:bg-primary/90 transition"
            >
              Subscribe
            </button>
          </form>
        </div>

        {/* Social Media */}
        <div>
          <h3 className="text-base font-semibold mb-2">Connect with Us</h3>
          <div className="flex space-x-4 mt-2">
            <a
              href="https://instagram.com/yourstudio"
              target="_blank"
              rel="noreferrer"
              className="text-xl hover:text-primary transition"
            >
              <FaInstagram />
            </a>
            <a
              href="https://linkedin.com/in/yourstudio"
              target="_blank"
              rel="noreferrer"
              className="text-xl hover:text-primary transition"
            >
              <FaLinkedin />
            </a>
            <a
              href="https://twitter.com/yourstudio"
              target="_blank"
              rel="noreferrer"
              className="text-xl hover:text-primary transition"
            >
              <FaTwitter />
            </a>
            <a
              href="https://github.com/yourstudio"
              target="_blank"
              rel="noreferrer"
              className="text-xl hover:text-primary transition"
            >
              <FaGithub />
            </a>
            <a
              href="mailto:info@fitbookstudio.com"
              className="text-xl hover:text-primary transition"
            >
              <FaEnvelope />
            </a>
          </div>
        </div>
      </div>

      <div className="py-6 text-center text-sm text-gray-500 border-t">
        <p>&copy; 2025 FitBook Studio. All rights reserved.</p>
      </div>
    </footer>
  );
};

export default Footer;
