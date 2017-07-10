class Parallel < Formula
  desc "Shell command parallelization utility"
  homepage "https://savannah.gnu.org/projects/parallel/"
  url "https://ftp.gnu.org/gnu/parallel/parallel-20170622.tar.bz2"
  mirror "https://ftpmirror.gnu.org/parallel/parallel-20170622.tar.bz2"
  sha256 "b3324ca1e7553b9903e0ab1fbb461334986eadfcd9b33e09013a6bac3c279645"
  head "https://git.savannah.gnu.org/git/parallel.git"

  conflicts_with "moreutils", :because => "both install a 'parallel' executable."

  def install
    system "./configure", "--prefix=#{prefix}"
    system "make", "install"
  end

  test do
    assert_equal "test\ntest\n",
                 shell_output("#{bin}/parallel --will-cite echo ::: test test")
  end
end
