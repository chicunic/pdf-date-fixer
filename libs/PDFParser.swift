import Foundation
import PDFKit

@_cdecl("parsePDF")
public func parsePDF(_ filePath: UnsafePointer<CChar>) -> UnsafePointer<CChar>? {
    guard let path = String(cString: filePath).addingPercentEncoding(withAllowedCharacters: .urlPathAllowed),
          let url = URL(string: "file://" + path),
          let pdfDocument = PDFDocument(url: url) else {
        return nil
    }

    var textContent = ""
    for pageIndex in 0..<pdfDocument.pageCount {
        if let page = pdfDocument.page(at: pageIndex) {
            textContent += page.string ?? ""
        }
    }

    // Convert to mutable pointer
    let mutablePtr = strdup(textContent)
    return mutablePtr.map { UnsafePointer($0) }
}

@_cdecl("freeString")
public func freeString(_ ptr: UnsafeMutablePointer<CChar>?) {
    if let ptr = ptr {
        free(ptr)
    }
}
